package service

import "time"

func (s *Service) CreateSlug(slugName string) (int, error) {
	return s.repository.CreateSlug(slugName)
}

func (s *Service) DeleteSlug(slugName string) error {
	return s.repository.DeleteSlug(slugName)
}

func (s *Service) UpdateUserSlugs(userId int, addSlugNames []string, deleteSlugNames []string, ttl uint64) error {
	if ttl == 0 {
		return s.repository.UpdateUserSlugs(userId, addSlugNames, deleteSlugNames, time.Time{})
	}
	validUntil := time.Now().Add(time.Duration(ttl) * time.Second)
	return s.repository.UpdateUserSlugs(userId, addSlugNames, deleteSlugNames, validUntil)
}

func (s *Service) GetUserSlugs(userId int) ([]string, error) {
	return s.repository.GetUserSlugs(userId)
}
