package service

import (
	"github.com/zpix1/avito-test-task/pkg/repository"
	"time"
)

type Service struct {
	repository repository.Implementation
}

type SlugsImplementation interface {
	CreateSlug(slugName string, autoAddPercent uint32) (int, error)
	DeleteSlug(slugName string) error
	UpdateUserSlugs(userId int, addSlugNames []string, deleteSlugNames []string, ttl uint64) error
	GetUserSlugs(userId int) ([]string, error)

	GetSlugHistoryCsv(userId int, startDate time.Time, endDate time.Time) (string, error)
}

type Implementation interface {
	SlugsImplementation
}

func NewService(repository repository.Implementation) Implementation {
	return &Service{repository: repository}
}
