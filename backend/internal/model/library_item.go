package model

type LibraryItem struct {
	ID         int    `json:"id"`
	UserID     int    `json:"user_id"`
	GameID     int    `json:"game_id"`
	AddedBy    string `json:"added_by"`
	Platform   string `json:"platform,omitempty"`
	PlatformID string `json:"platform_id,omitempty"`
	Created_At string `json:"created_at"`
	Updated_At string `json:"updated_at"`
}
