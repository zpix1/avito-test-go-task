package postgres

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
	"time"
)

func (r *Repository) CreateSlug(slugName string) (int, error) {
	row := r.pool.QueryRow(context.Background(), "INSERT INTO slugs (name) VALUES ($1) RETURNING id", slugName)

	var slugId int
	if err := row.Scan(&slugId); err != nil {
		return 0, err
	}
	return slugId, nil
}

func (r *Repository) DeleteSlug(slugName string) error {
	rows, err := r.pool.Query(
		context.Background(),
		"SELECT user_id FROM slugs_users WHERE slug_name=$1 AND (valid_until >= CURRENT_TIMESTAMP OR valid_until IS NULL)",
		slugName,
	)

	for rows.Next() {
		var userId int
		err := rows.Scan(&userId)

		if err != nil {
			return err
		}

		go func() {
			err := r.SaveSlugActionHistory(userId, slugName, true)
			if err != nil {
				logrus.Warn("failed to add entry to slug history", slugName)
			}
		}()
	}
	_, err = r.pool.Exec(context.Background(), "DELETE FROM slugs WHERE name=$1", slugName)
	return err
}

func (r *Repository) UpdateUserSlugs(userId int, addSlugNames []string, deleteSlugNames []string, validUntil time.Time) error {
	batch := &pgx.Batch{}
	for _, addSlugName := range addSlugNames {
		if validUntil.IsZero() {
			batch.Queue("INSERT INTO slugs_users(slug_name, user_id) VALUES ($1, $2)", addSlugName, userId)
		} else {
			batch.Queue("INSERT INTO slugs_users(slug_name, user_id, valid_until) VALUES ($1, $2, $3)", addSlugName, userId, validUntil)
			go func() {
				err := r.SaveSlugActionHistoryWithTime(userId, addSlugName, true, validUntil)
				if err != nil {
					logrus.Warn("failed to add entry to slug history", addSlugName)
				}
			}()
		}
	}
	for _, deleteSlugName := range deleteSlugNames {
		// do not delete slugs that are not valid anymore, they are handled automatically
		batch.Queue("DELETE FROM slugs_users WHERE slug_name=$1 AND user_id=$2 AND valid_until >= CURRENT_TIMESTAMP", deleteSlugName, userId)
	}
	br := r.pool.SendBatch(context.Background(), batch)
	_, err := br.Exec()
	if err != nil {
		return err
	}
	go func() {
		for _, addSlugName := range addSlugNames {
			err := r.SaveSlugActionHistory(userId, addSlugName, false)
			// history is not so important to stop slugs requests
			if err != nil {
				logrus.Warn("failed to add entry to slug history", addSlugName)
			}
		}
		for _, deleteSlugName := range deleteSlugNames {
			err := r.SaveSlugActionHistory(userId, deleteSlugName, true)
			if err != nil {
				logrus.Warn("failed to add entry to slug history", deleteSlugName)
			}
		}
	}()
	return nil
}

func (r *Repository) GetUserSlugs(userId int) ([]string, error) {
	rows, err := r.pool.Query(
		context.Background(),
		"SELECT slug_name FROM slugs_users WHERE user_id=$1 AND (valid_until >= CURRENT_TIMESTAMP OR valid_until IS NULL)",
		userId,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var slugNames []string
	for rows.Next() {
		var sn string
		err := rows.Scan(&sn)

		if err != nil {
			return nil, err
		}
		slugNames = append(slugNames, sn)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return slugNames, nil
}
