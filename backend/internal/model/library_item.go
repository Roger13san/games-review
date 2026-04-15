package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LibraryItem struct {
	ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID primitive.ObjectID `json:"user_id" bson:"user_id"`
	// GameID é o appid da Steam — identificador único e universal do jogo.
	// Jogos adicionados manualmente (fora da Steam) usam 0.
	GameID     uint32    `json:"game_id" bson:"game_id"`
	AddedBy    string    `json:"added_by" bson:"added_by"`
	Platform   string    `json:"platform,omitempty" bson:"platform,omitempty"`
	PlatformID string    `json:"platform_id,omitempty" bson:"platform_id,omitempty"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" bson:"updated_at"`
}
