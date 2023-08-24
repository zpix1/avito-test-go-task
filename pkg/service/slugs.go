package service

func (s *Service) CreateSlug(slugName string) (int, error) {
	return s.repository.CreateSlug(slugName)
}

func (s *Service) DeleteSlug(slugName string) error {
	return s.repository.DeleteSlug(slugName)
}

func (s *Service) UpdateUserSlugs(userId int, addSlugNames []string, deleteSlugNames []string) error {
	return s.repository.UpdateUserSlugs(userId, addSlugNames, deleteSlugNames)
}

func (s *Service) GetUserSlugs(userId int) ([]string, error) {
	return s.repository.GetUserSlugs(userId)
}
