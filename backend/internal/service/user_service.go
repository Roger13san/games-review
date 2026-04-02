package service

import (
	"errors"
	"strings"
	"time"

	"github.com/Roger13san/games-review/backend/internal/model"
	"github.com/Roger13san/games-review/backend/internal/repository"
	"github.com/Roger13san/games-review/backend/internal/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RegisterUser(user model.User) (model.User, error) {
	if err := validateUser(user); err != nil {
		return model.User{}, err
	}
	user.ID = primitive.NewObjectID()
	user.Password = util.HashPassword(user.Password)
	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt
	return repository.CreateUser(user)
}

func LoginUser(email, password string) (model.User, error) {
	user, err := repository.GetUserByEmail(email)
	if err != nil {
		return model.User{}, errors.New("usuário não encontrado")
	}
	if !util.CheckPasswordHash(password, user.Password) {
		return model.User{}, errors.New("senha inválida")
	}
	user.LastLoginAt = time.Now()
	return user, nil
}

func GetUserByID(id primitive.ObjectID) (model.User, error) {
	return repository.GetUserByID(id)
}

func ListUsers() ([]model.User, error) {
	return repository.ListUsers()
}

func UpdateUser(id primitive.ObjectID, user model.User) (model.User, error) {
	user.UpdatedAt = time.Now()
	return repository.UpdateUser(id, user)
}

func DeleteUser(id primitive.ObjectID) error {
	return repository.DeleteUser(id)
}

func validateUser(user model.User) error {
	if strings.TrimSpace(user.Username) == "" {
		return errors.New("username é obrigatório")
	}
	if strings.TrimSpace(user.Email) == "" {
		return errors.New("email é obrigatório")
	}
	if strings.TrimSpace(user.Password) == "" {
		return errors.New("senha é obrigatória")
	}
	return nil
}
