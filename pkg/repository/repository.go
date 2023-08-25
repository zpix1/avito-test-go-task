package repository

import (
	entity "github.com/zpix1/avito-test-task/pkg/entities"
	"time"
)

type SlugsImplementation interface {
	CreateSlug(slugName string) (int, error)
	DeleteSlug(slugName string) error
	UpdateUserSlugs(userId int, addSlugNames []string, deleteSlugNames []string) error
	GetUserSlugs(userId int) ([]string, error)

	GetSlugHistory(userId int, startDate time.Time, endDate time.Time) ([]entity.SlugHistoryEntry, error)
	SaveSlugActionHistory(userId int, slugName string, removed bool) error
}

type Implementation interface {
	SlugsImplementation
}
