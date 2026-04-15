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

func GetReviews() ([]model.Review, error) {
	collection := database.GetCollection("reviews")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	findOptions := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})
	cursor, err := collection.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		return nil, err
	}

	var reviews []model.Review
	if err = cursor.All(ctx, &reviews); err != nil {
		return nil, err
	}
	return reviews, nil
}

// GetReviewsByGameID retorna todas as reviews de um jogo específico pelo appid da Steam.
func GetReviewsByGameID(gameID uint32) ([]model.Review, error) {
	collection := database.GetCollection("reviews")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	findOptions := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})
	cursor, err := collection.Find(ctx, bson.M{"game_id": gameID}, findOptions)
	if err != nil {
		return nil, err
	}

	var reviews []model.Review
	if err = cursor.All(ctx, &reviews); err != nil {
		return nil, err
	}
	return reviews, nil
}

// GetReviewsByUserID retorna todas as reviews escritas por um usuário específico.
func GetReviewsByUserID(userID primitive.ObjectID) ([]model.Review, error) {
	collection := database.GetCollection("reviews")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	findOptions := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})
	cursor, err := collection.Find(ctx, bson.M{"user_id": userID}, findOptions)
	if err != nil {
		return nil, err
	}

	var reviews []model.Review
	if err = cursor.All(ctx, &reviews); err != nil {
		return nil, err
	}
	return reviews, nil
}

func GetReviewByID(id primitive.ObjectID) (model.Review, error) {
	collection := database.GetCollection("reviews")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var review model.Review
	err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&review)
	if err != nil {
		return model.Review{}, err
	}

	return review, nil
}

func CreateReview(review model.Review) (model.Review, error) {
	collection := database.GetCollection("reviews")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	review.ID = primitive.NewObjectID()
	review.CreatedAt = time.Now()
	review.UpdatedAt = time.Now()

	_, err := collection.InsertOne(ctx, review)
	if err != nil {
		return model.Review{}, err
	}

	return review, nil
}

func UpdateReview(id primitive.ObjectID, review model.Review) (model.Review, error) {
	collection := database.GetCollection("reviews")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	review.UpdatedAt = time.Now()

	_, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": review},
	)
	if err != nil {
		return model.Review{}, err
	}

	return review, nil
}

func DeleteReview(id primitive.ObjectID) error {
	collection := database.GetCollection("reviews")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
