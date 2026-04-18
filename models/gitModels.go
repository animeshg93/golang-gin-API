package models

type GitUser struct {
	Login      string `json:"login"`
	Email      string `json:"email"`
	Blog       string `json:"blog"`
	GravitarID string `json:"gravatar_id"`
	Location   string `json:"location"`
}

type GitRepo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	HTMLURL     string `json:"html_url"`
}
