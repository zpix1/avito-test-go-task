package repository

import (
	entity "github.com/zpix1/avito-test-task/pkg/entities"
	"time"
)

type SlugsImplementation interface {
	CreateSlug(slugName string, autoAddWeight uint32) (int, error)
	DeleteSlug(slugName string) error
	UpdateUserSlugs(userId int, addSlugNames []string, deleteSlugNames []string, validUntil time.Time) error
	GetUserSlugs(userId int) ([]string, error)

	GetSlugHistory(userId int, startDate time.Time, endDate time.Time) ([]entity.SlugHistoryEntry, error)
	SaveSlugActionHistory(userId int, slugName string, removed bool) error

	GetAutoAddSlugs(userId int) ([]entity.SlugAutoAdd, error)
}

type Implementation interface {
	SlugsImplementation
}
