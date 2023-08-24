package entities

type Slug struct {
	Name string `json:"name"`
}

type SlugUpdate struct {
	UserId          int      `json:"user_id"`
	AddSlugNames    []string `json:"add_slug_names"`
	DeleteSlugNames []string `json:"delete_slug_names"`
}

type GetSlugsResponse struct {
	SlugNames []string `json:"slug_names"`
}