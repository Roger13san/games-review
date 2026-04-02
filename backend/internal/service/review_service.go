package service

import (
	"errors"
	"strings"

	"github.com/Roger13san/games-review/backend/internal/model"
	"github.com/Roger13san/games-review/backend/internal/repository"
)

func GetReviewsService() ([]model.Review, error) {
	return repository.GetReviews()
}

func GetReviewByIDService(id int) (model.Review, error) {
	if id <= 0 {
		return model.Review{}, errors.New("id inválido")
	}

	return repository.GetReviewByID(id)
}

func CreateReviewService(review model.Review) (model.Review, error) {
	if err := validateReview(review); err != nil {
		return model.Review{}, err
	}

	return repository.CreateReview(review)
}

func UpdateReviewService(id int, review model.Review) (model.Review, error) {
	if id <= 0 {
		return model.Review{}, errors.New("id inválido")
	}

	if err := validateReview(review); err != nil {
		return model.Review{}, err
	}

	return repository.UpdateReview(id, review)
}

func DeleteReviewService(id int) error {
	if id <= 0 {
		return errors.New("id inválido")
	}

	return repository.DeleteReview(id)
}

func validateReview(review model.Review) error {
	if strings.TrimSpace(review.Title) == "" {
		return errors.New("título é obrigatório")
	}

	if strings.TrimSpace(review.Content) == "" {
		return errors.New("conteúdo é obrigatório")
	}

	if review.GameID <= 0 {
		return errors.New("game_id inválido")
	}

	if review.UserID <= 0 {
		return errors.New("user_id inválido")
	}

	if review.Rating < 1 || review.Rating > 5 {
		return errors.New("rating deve estar entre 1 e 5")
	}

	return nil
}
