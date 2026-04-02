package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/Roger13san/games-review/backend/internal/middleware"
	"github.com/Roger13san/games-review/backend/internal/model"
	"github.com/Roger13san/games-review/backend/internal/service"
)

func HandleReviews(w http.ResponseWriter, r *http.Request) {
	id, hasID, err := extractIDFromPath(r.URL.Path, "/reviews")
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

func listarReviews(w http.ResponseWriter, r *http.Request) {
	reviews, err := service.GetReviewsService()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reviews)
}

func listarReviewByID(w http.ResponseWriter, id int) {
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
	err := json.NewDecoder(r.Body).Decode(&novaReview)
	if err != nil {
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

func atualizarReview(w http.ResponseWriter, r *http.Request, id int) {
	var review model.Review
	err := json.NewDecoder(r.Body).Decode(&review)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Dados inválidos"))
		return
	}

	updatedReview, err := service.UpdateReviewService(id, review)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedReview)
}

func deletarReview(w http.ResponseWriter, id int) {
	err := service.DeleteReviewService(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func extractIDFromPath(path string, basePath string) (int, bool, error) {
	trimmedPath := strings.TrimSuffix(path, "/")
	trimmedBasePath := strings.TrimSuffix(basePath, "/")

	if trimmedPath == trimmedBasePath {
		return 0, false, nil
	}

	prefix := trimmedBasePath + "/"
	if !strings.HasPrefix(trimmedPath, prefix) {
		return 0, false, nil
	}

	idPart := strings.TrimPrefix(trimmedPath, prefix)
	if strings.Contains(idPart, "/") || idPart == "" {
		return 0, false, strconv.ErrSyntax
	}

	id, err := strconv.Atoi(idPart)
	if err != nil {
		return 0, false, err
	}

	return id, true, nil
}
