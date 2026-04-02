package model

type Game struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	CoverURL   string `json:"cover_url,omitempty"`
	Created_At string `json:"created_at"`
	Updated_At string `json:"updated_at"`
}
