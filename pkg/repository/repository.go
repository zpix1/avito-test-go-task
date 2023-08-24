package repository

type SlugsImplementation interface {
	CreateSlug(slugName string) (int, error)
	DeleteSlug(slugName string) error
	UpdateUserSlugs(userId int, addSlugNames []string, deleteSlugNames []string) error
	GetUserSlugs(userId int) ([]string, error)
}

type Implementation interface {
	SlugsImplementation
}
