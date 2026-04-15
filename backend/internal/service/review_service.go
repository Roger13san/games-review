package service

import (
	"errors"
	"strings"

	"github.com/Roger13san/games-review/backend/internal/model"
	"github.com/Roger13san/games-review/backend/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetReviewsService() ([]model.Review, error) {
	return repository.GetReviews()
}

func GetReviewsByGameIDService(gameID uint32) ([]model.Review, error) {
	return repository.GetReviewsByGameID(gameID)
}

func GetReviewsByUserIDService(userID primitive.ObjectID) ([]model.Review, error) {
	return repository.GetReviewsByUserID(userID)
}

func GetReviewByIDService(id primitive.ObjectID) (model.Review, error) {
	return repository.GetReviewByID(id)
}

func CreateReviewService(review model.Review) (model.Review, error) {
	if err := validateReview(review); err != nil {
		return model.Review{}, err
	}

	return repository.CreateReview(review)
}

func UpdateReviewService(id primitive.ObjectID, review model.Review) (model.Review, error) {
	if err := validateReview(review); err != nil {
		return model.Review{}, err
	}

	return repository.UpdateReview(id, review)
}

func DeleteReviewService(id primitive.ObjectID) error {
	return repository.DeleteReview(id)
}

func validateReview(review model.Review) error {
	if strings.TrimSpace(review.Title) == "" {
		return errors.New("título é obrigatório")
	}

	if strings.TrimSpace(review.Content) == "" {
		return errors.New("conteúdo é obrigatório")
	}

	// GameID é o appid da Steam — qualquer valor maior que 0 é válido.
	if review.GameID == 0 {
		return errors.New("game_id inválido")
	}

	if review.UserID.IsZero() {
		return errors.New("user_id inválido")
	}

	if review.Rating < 1 || review.Rating > 5 {
		return errors.New("rating deve estar entre 1 e 5")
	}

	return nil
}
