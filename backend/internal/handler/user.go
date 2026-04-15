package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Roger13san/games-review/backend/internal/middleware"
	"github.com/Roger13san/games-review/backend/internal/model"
	"github.com/Roger13san/games-review/backend/internal/service"
	"github.com/Roger13san/games-review/backend/internal/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func HandleUsers(w http.ResponseWriter, r *http.Request) {
	id, hasID, _ := extractObjectIDFromPath(r.URL.Path, "/users")

	switch r.Method {
	case http.MethodGet:
		if hasID {
			getUserByID(w, id)
			return
		}
		listUsers(w)
	case http.MethodPost:
		if strings.HasSuffix(r.URL.Path, "/login") {
			loginUser(w, r)
			return
		}
		registerUser(w, r)
	case http.MethodPut:
		if !hasID {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("ID obrigatório"))
			return
		}
		if _, authErr := middleware.RequireAuth(r); authErr != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(authErr.Error()))
			return
		}
		updateUser(w, r, id)
	case http.MethodDelete:
		if !hasID {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("ID obrigatório"))
			return
		}
		if _, authErr := middleware.RequireAuth(r); authErr != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(authErr.Error()))
			return
		}
		deleteUser(w, id)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Método não permitido"))
	}
}

func registerUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Dados inválidos"))
		return
	}
	created, err := service.RegisterUser(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Dados inválidos"))
		return
	}
	user, err := service.LoginUser(req.Email, req.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}
	token, err := util.GenerateJWT(user.ID.Hex())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Erro ao gerar token"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func getUserByID(w http.ResponseWriter, id primitive.ObjectID) {
	user, err := service.GetUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Usuário não encontrado"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func listUsers(w http.ResponseWriter) {
	users, err := service.ListUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func updateUser(w http.ResponseWriter, r *http.Request, id primitive.ObjectID) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Dados inválidos"))
		return
	}
	updated, err := service.UpdateUser(id, user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updated)
}

func deleteUser(w http.ResponseWriter, id primitive.ObjectID) {
	if err := service.DeleteUser(id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

