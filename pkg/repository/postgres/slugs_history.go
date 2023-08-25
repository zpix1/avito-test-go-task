package postgres

import (
	"context"
	entity "github.com/zpix1/avito-test-task/pkg/entities"
	"time"
)

func (r *Repository) SaveSlugActionHistory(userId int, slugName string, removed bool) error {
	_, err := r.pool.Exec(
		context.Background(),
		"INSERT INTO slugs_history(user_id, slug_name, removed) VALUES ($1, $2, $3)",
		userId, slugName, removed)
	return err
}

func (r *Repository) GetSlugHistory(userId int, startDate time.Time, endDate time.Time) ([]entity.SlugHistoryEntry, error) {
	rows, err := r.pool.Query(context.Background(),
		"SELECT user_id, slug_name, removed, created_at FROM slugs_history WHERE user_id=$1 AND $2 <= created_at  AND created_at<= $3",
		userId, startDate, endDate)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var slugHistory []entity.SlugHistoryEntry
	for rows.Next() {
		var slugHistoryEntry entity.SlugHistoryEntry
		err := rows.Scan(&slugHistoryEntry.UserId,
			&slugHistoryEntry.SlugName,
			&slugHistoryEntry.Removed,
			&slugHistoryEntry.DoneAt)

		if err != nil {
			return nil, err
		}
		slugHistory = append(slugHistory, slugHistoryEntry)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return slugHistory, nil
}
