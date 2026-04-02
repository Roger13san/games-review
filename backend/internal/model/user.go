package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	SteamID     string             `bson:"steam_id,omitempty" json:"steam_id,omitempty"`
	Username    string             `bson:"username" json:"username"`
	Email       string             `bson:"email" json:"email"`
	Password    string             `bson:"password" json:"password,omitempty"`
	AvatarURL   string             `bson:"avatar_url,omitempty" json:"avatar_url,omitempty"`
	ProfileURL  string             `bson:"profile_url,omitempty" json:"profile_url,omitempty"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
	LastLoginAt time.Time          `bson:"last_login_at,omitempty" json:"last_login_at,omitempty"`
}
