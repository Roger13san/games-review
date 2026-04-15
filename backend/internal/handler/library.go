package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Roger13san/games-review/backend/internal/middleware"
	"github.com/Roger13san/games-review/backend/internal/model"
	"github.com/Roger13san/games-review/backend/internal/repository"
	"github.com/Roger13san/games-review/backend/internal/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// HandleLibrary roteia GET e POST em /library.
func HandleLibrary(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getLibrary(w, r)
	case http.MethodPost:
		addToLibrary(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Método não permitido"))
	}
}

// HandleImportLibrary importa toda a biblioteca Steam do usuário de uma vez.
//
// Rota: POST /library/import
func HandleImportLibrary(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Método não permitido"))
		return
	}

	claims, err := middleware.RequireAuth(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	userOID, err := primitive.ObjectIDFromHex(claims.UserID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("userID inválido"))
		return
	}

	user, err := repository.GetUserByID(userOID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	games, err := service.GetOwnedGames(user.SteamID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	for _, game := range games {
		item := model.LibraryItem{
			UserID:     userOID,
			GameID:     game.AppID,
			AddedBy:    "steam",
			Platform:   "steam",
			PlatformID: fmt.Sprintf("%d", game.AppID),
		}
		if err := repository.UpsertLibraryItem(item); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"imported": len(games)})
}

// getLibrary retorna a biblioteca de jogos do usuário logado.
//
// Rota: GET /library
func getLibrary(w http.ResponseWriter, r *http.Request) {
	claims, err := middleware.RequireAuth(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	userOID, err := primitive.ObjectIDFromHex(claims.UserID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("userID inválido"))
		return
	}

	items, err := repository.GetLibraryItems(userOID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if items == nil {
		items = []model.LibraryItem{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// addToLibrary adiciona um jogo manualmente à biblioteca do usuário.
// Usado para jogos jogados fora da Steam (Epic, GOG etc.) que têm AppID na Steam.
//
// Rota: POST /library
// Body: { "game_id": 730 }
func addToLibrary(w http.ResponseWriter, r *http.Request) {
	claims, err := middleware.RequireAuth(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	userOID, err := primitive.ObjectIDFromHex(claims.UserID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("userID inválido"))
		return
	}

	var body struct {
		GameID uint32 `json:"game_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.GameID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("game_id obrigatório e deve ser um appid Steam válido"))
		return
	}

	item := model.LibraryItem{
		UserID:     userOID,
		GameID:     body.GameID,
		AddedBy:    "manual",
		Platform:   "other",
		PlatformID: fmt.Sprintf("%d", body.GameID),
	}

	if err := repository.UpsertLibraryItem(item); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

