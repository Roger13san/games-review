package model

type Review struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	GameID     int    `json:"game_id"`
	UserID     int    `json:"user_id"`
	Rating     int    `json:"rating"`
	Created_At string `json:"created_at"`
	Updated_At string `json:"updated_at"`
}
