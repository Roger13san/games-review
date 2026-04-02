package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/Roger13san/games-review/backend/internal/util"
)

func RequireAuth(r *http.Request) (*util.Claims, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return nil, errors.New("token ausente")
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
		return nil, errors.New("formato de token inválido")
	}

	claims, err := util.ParseJWT(parts[1])
	if err != nil {
		return nil, errors.New("token inválido")
	}

	return claims, nil
}
