package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Game struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title     string             `json:"title" bson:"title"`
	CoverURL  string             `json:"cover_url,omitempty" bson:"cover_url,omitempty"`
	// SteamAppID é o appid da Steam — presente quando o jogo veio da Steam,
	// zero quando foi adicionado manualmente.
	SteamAppID uint32            `json:"steam_app_id,omitempty" bson:"steam_app_id,omitempty"`
	CreatedAt  time.Time         `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at" bson:"updated_at"`
}
