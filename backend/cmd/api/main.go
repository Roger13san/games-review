package main

import (
	"log"
	"net/http"

	"github.com/Roger13san/games-review/backend/internal/database"
	"github.com/Roger13san/games-review/backend/internal/router"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar .env")
	}

	err = database.ConnectMongo()
	if err != nil {
		log.Fatal("Erro ao conectar no MongoDB:", err)
	}

	router.RegisterRoutes()
	log.Fatal(http.ListenAndServe(":8081", nil))
}
