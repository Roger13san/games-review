package router

import (
	"net/http"

	"github.com/Roger13san/games-review/backend/internal/handler"
)

func RegisterRoutes() {
	http.HandleFunc("/", handler.HandleRoot)

	// Auth Steam
	http.HandleFunc("/auth/steam", handler.HandleSteamLogin)
	http.HandleFunc("/auth/steam/callback", handler.HandleSteamCallback)

	// Reviews
	http.HandleFunc("/reviews", handler.HandleReviews)
	http.HandleFunc("/reviews/", handler.HandleReviews)

	// Users
	http.HandleFunc("/users", handler.HandleUsers)
	http.HandleFunc("/users/", handler.HandleUsers)
	// Library
	http.HandleFunc("/library", handler.HandleLibrary)
	http.HandleFunc("/library/import", handler.HandleImportLibrary)
}
