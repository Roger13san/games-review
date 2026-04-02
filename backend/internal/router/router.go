package router

import (
	"net/http"

	"github.com/Roger13san/games-review/backend/internal/handler"
)

func RegisterRoutes() {
	http.HandleFunc("/", handler.HandleRoot)
	http.HandleFunc("/reviews", handler.HandleReviews)
	http.HandleFunc("/reviews/", handler.HandleReviews)
	http.HandleFunc("/users", handler.HandleUsers)
	http.HandleFunc("/users/", handler.HandleUsers)
}
