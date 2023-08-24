package postgres

import (
	"context"
	"github.com/jackc/pgx/v4"
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
	_, err := r.pool.Exec(context.Background(), "DELETE FROM slugs WHERE name=$1", slugName)
	return err
}

func (r *Repository) UpdateUserSlugs(userId int, addSlugNames []string, deleteSlugNames []string) error {
	batch := &pgx.Batch{}
	for _, addSlugName := range addSlugNames {
		batch.Queue("INSERT INTO slugs_users(slug_name, user_id) VALUES ($1, $2)", addSlugName, userId)
	}
	for _, deleteSlugName := range deleteSlugNames {
		batch.Queue("DELETE FROM slugs_users WHERE slug_name=$1 AND user_id=$2", deleteSlugName, userId)
	}
	br := r.pool.SendBatch(context.Background(), batch)
	_, err := br.Exec()
	return err
}

func (r *Repository) GetUserSlugs(userId int) ([]string, error) {
	rows, err := r.pool.Query(context.Background(), "SELECT slug_name FROM slugs_users WHERE user_id=$1", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var slugNames []string
	for rows.Next() {
		var r string
		err := rows.Scan(&r)

		if err != nil {
			return nil, err
		}
		slugNames = append(slugNames, r)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return slugNames, nil
}
