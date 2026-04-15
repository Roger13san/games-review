package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Roger13san/games-review/backend/internal/middleware"
	"github.com/Roger13san/games-review/backend/internal/model"
	"github.com/Roger13san/games-review/backend/internal/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func HandleReviews(w http.ResponseWriter, r *http.Request) {
	id, hasID, err := extractObjectIDFromPath(r.URL.Path, "/reviews")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID inválido"))
		return
	}

	switch r.Method {
	case http.MethodGet:
		if hasID {
			listarReviewByID(w, id)
			return
		}
		listarReviews(w, r)
	case http.MethodPost:
		if hasID {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Método não permitido"))
			return
		}
		if _, authErr := middleware.RequireAuth(r); authErr != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(authErr.Error()))
			return
		}
		criarReview(w, r)
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
		atualizarReview(w, r, id)
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
		deletarReview(w, id)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Método não permitido"))
	}
}

// listarReviews lida com GET /reviews, suportando dois filtros via query params:
//   - ?game_id=730       → reviews de um jogo específico (appid Steam)
//   - ?user_id=<objectid> → reviews de um usuário específico
//
// Sem filtros, retorna todas as reviews.
func listarReviews(w http.ResponseWriter, r *http.Request) {
	var reviews []model.Review
	var err error

	if gameIDStr := r.URL.Query().Get("game_id"); gameIDStr != "" {
		// Converte o appid (uint32) recebido como string na query.
		gameID64, parseErr := strconv.ParseUint(gameIDStr, 10, 32)
		if parseErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("game_id inválido"))
			return
		}
		reviews, err = service.GetReviewsByGameIDService(uint32(gameID64))

	} else if userIDStr := r.URL.Query().Get("user_id"); userIDStr != "" {
		userOID, parseErr := primitive.ObjectIDFromHex(userIDStr)
		if parseErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("user_id inválido"))
			return
		}
		reviews, err = service.GetReviewsByUserIDService(userOID)

	} else {
		reviews, err = service.GetReviewsService()
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if reviews == nil {
		reviews = []model.Review{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reviews)
}

func listarReviewByID(w http.ResponseWriter, id primitive.ObjectID) {
	review, err := service.GetReviewByIDService(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(review)
}

func criarReview(w http.ResponseWriter, r *http.Request) {
	var novaReview model.Review
	if err := json.NewDecoder(r.Body).Decode(&novaReview); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Dados inválidos"))
		return
	}

	review, err := service.CreateReviewService(novaReview)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(review)
}

func atualizarReview(w http.ResponseWriter, r *http.Request, id primitive.ObjectID) {
	var review model.Review
	if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Dados inválidos"))
		return
	}

	updated, err := service.UpdateReviewService(id, review)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updated)
}

func deletarReview(w http.ResponseWriter, id primitive.ObjectID) {
	if err := service.DeleteReviewService(id); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

