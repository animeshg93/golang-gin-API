package models

type GitUser struct {
	Login      string `json:"login"`
	Email      string `json:"email"`
	Blog       string `json:"blog"`
	GravitarID string `json:"gravatar_id"`
	Location   string `json:"location"`
}
