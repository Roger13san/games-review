package repository

import (
	"context"
	"time"

	"github.com/Roger13san/games-review/backend/internal/database"
	"github.com/Roger13san/games-review/backend/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateUser(user model.User) (model.User, error) {
	collection := database.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := collection.InsertOne(ctx, user)
	if err == nil {
		user.ID = result.InsertedID.(primitive.ObjectID)
	}
	return user, err
}

func GetUserByID(id primitive.ObjectID) (model.User, error) {
	collection := database.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var user model.User
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	return user, err
}

func GetUserByEmail(email string) (model.User, error) {
	collection := database.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var user model.User
	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	return user, err
}

func UpdateUser(id primitive.ObjectID, user model.User) (model.User, error) {
	collection := database.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	user.UpdatedAt = time.Now()
	_, err := collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": user})
	return user, err
}

func DeleteUser(id primitive.ObjectID) error {
	collection := database.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func ListUsers() ([]model.User, error) {
	collection := database.GetCollection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.D{}, options.Find())
	if err != nil {
		return nil, err
	}
	var users []model.User
	err = cursor.All(ctx, &users)
	return users, err
}
