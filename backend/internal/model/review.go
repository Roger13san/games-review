package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Review struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title     string             `json:"title" bson:"title"`
	Content   string             `json:"content" bson:"content"`
	// GameID é o appid da Steam — identificador universal do jogo.
	// Não precisa existir no banco local, a Steam é a fonte de verdade.
	GameID    uint32             `json:"game_id" bson:"game_id"`
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`
	Rating    int                `json:"rating" bson:"rating"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}
