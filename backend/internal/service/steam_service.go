package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Roger13san/games-review/backend/internal/model"
)

// steamClient é um http.Client com timeout configurado.
// Sem timeout, se a Steam API demorar ou travar a requisição
// fica pendurada indefinidamente e prende a goroutine.
var steamClient = &http.Client{
	Timeout: 10 * time.Second,
}

// SteamProfile representa os dados de perfil que a Steam retorna
// no endpoint GetPlayerSummaries.
type SteamProfile struct {
	SteamID    string `json:"steamid"`
	Username   string `json:"personaname"` // nome de exibição do perfil
	AvatarURL  string `json:"avatarfull"`  // URL do avatar em tamanho grande
	ProfileURL string `json:"profileurl"`  // link pro perfil público na Steam
}

// steamPlayerSummaryResponse reflete exatamente o JSON que a Steam retorna.
// A estrutura é: { "response": { "players": [ { ...perfil... } ] } }
type steamPlayerSummaryResponse struct {
	Response struct {
		Players []SteamProfile `json:"players"`
	} `json:"response"`
}

// steamOwnedGamesResponse reflete o JSON retornado pelo endpoint GetOwnedGames.
// A estrutura é: { "response": { "game_count": N, "games": [ {...}, ... ] } }
type steamOwnedGamesResponse struct {
	Response struct {
		GameCount int                `json:"game_count"`
		Games     []model.SteamGame  `json:"games"`
	} `json:"response"`
}

// GetSteamProfile busca o perfil público de um usuário na Steam Web API
// usando o steamid64 obtido após o login OpenID.
//
// Endpoint: GET /ISteamUser/GetPlayerSummaries/v2/
func GetSteamProfile(steamID string) (SteamProfile, error) {
	apiKey := os.Getenv("STEAM_API_KEY")
	if apiKey == "" {
		return SteamProfile{}, fmt.Errorf("STEAM_API_KEY não definida no .env")
	}

	url := fmt.Sprintf(
		"https://api.steampowered.com/ISteamUser/GetPlayerSummaries/v2/?key=%s&steamids=%s",
		apiKey,
		steamID,
	)

	resp, err := steamClient.Get(url)
	if err != nil {
		return SteamProfile{}, fmt.Errorf("erro ao chamar Steam API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return SteamProfile{}, fmt.Errorf("Steam API retornou status %d", resp.StatusCode)
	}

	var result steamPlayerSummaryResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return SteamProfile{}, fmt.Errorf("erro ao decodificar resposta da Steam: %w", err)
	}

	if len(result.Response.Players) == 0 {
		return SteamProfile{}, fmt.Errorf("nenhum perfil encontrado para steamid %s", steamID)
	}

	return result.Response.Players[0], nil
}

// GetOwnedGames busca todos os jogos da biblioteca de um usuário na Steam.
//
// Endpoint: GET /IPlayerService/GetOwnedGames/v1/
//
// Parâmetros usados:
//   - include_appinfo=true      → traz nome, ícone e logo do jogo
//   - include_played_free_games → inclui jogos gratuitos que já foram jogados
//
// Retorna slice de SteamGame com todos os campos mapeados da documentação.
// Se a biblioteca for privada ou o steamid inválido, retorna slice vazio.
func GetOwnedGames(steamID string) ([]model.SteamGame, error) {
	apiKey := os.Getenv("STEAM_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("STEAM_API_KEY não definida no .env")
	}

	url := fmt.Sprintf(
		"https://api.steampowered.com/IPlayerService/GetOwnedGames/v1/?key=%s&steamid=%s&include_appinfo=true&include_played_free_games=true",
		apiKey,
		steamID,
	)

	resp, err := steamClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("erro ao chamar Steam API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Steam API retornou status %d", resp.StatusCode)
	}

	var result steamOwnedGamesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("erro ao decodificar resposta da Steam: %w", err)
	}

	// Biblioteca privada ou sem jogos — não é erro, apenas retorna vazio.
	if result.Response.GameCount == 0 {
		return []model.SteamGame{}, nil
	}

	return result.Response.Games, nil
}
