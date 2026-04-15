package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Roger13san/games-review/backend/internal/model"
	"github.com/Roger13san/games-review/backend/internal/repository"
	"github.com/Roger13san/games-review/backend/internal/service"
	"github.com/Roger13san/games-review/backend/internal/util"
	"github.com/yohcop/openid-go"
)

// nonceStore e discoveryCache são exigidos pela lib openid-go.
// Em produção você trocaria por implementações persistentes
// (Redis, banco etc.), mas para desenvolvimento o in-memory funciona bem.
var (
	nonceStore     = openid.NewSimpleNonceStore()
	discoveryCache = openid.NewSimpleDiscoveryCache()
)

// HandleSteamLogin inicia o fluxo de autenticação via Steam OpenID 2.0.
//
// O que acontece aqui:
//  1. Montamos a URL do endpoint OpenID da Steam.
//  2. Dizemos à Steam qual é a nossa URL de callback (return_to).
//  3. Redirecionamos o usuário para a Steam — ele vai ver a tela
//     de "Entrar com Steam" que ele já conhece.
//
// Rota: GET /auth/steam
func HandleSteamLogin(w http.ResponseWriter, r *http.Request) {
	baseURL := os.Getenv("APP_BASE_URL")
	if baseURL == "" {
		log.Println("[steam] APP_BASE_URL não definida no .env")
		http.Error(w, "Configuração interna inválida", http.StatusInternalServerError)
		return
	}

	// return_to: URL do nosso callback — a Steam vai redirecionar o usuário pra cá
	// após o login bem-sucedido.
	callbackURL := fmt.Sprintf("%s/auth/steam/callback", baseURL)

	// realm: domínio raiz do nosso site. A Steam mostra isso pro usuário
	// na tela de autorização ("você está entrando em...").
	realm := baseURL

	// openid.RedirectURL monta a URL completa do endpoint OpenID da Steam
	// com todos os parâmetros necessários (ns, mode, return_to, realm etc.)
	redirectURL, err := openid.RedirectURL(
		"https://steamcommunity.com/openid", // endpoint fixo da Steam
		callbackURL,
		realm,
	)
	if err != nil {
		log.Printf("[steam] erro ao montar redirect URL: %v", err)
		http.Error(w, "Erro ao iniciar autenticação", http.StatusInternalServerError)
		return
	}

	// Redireciona o usuário para a página de login da Steam.
	http.Redirect(w, r, redirectURL, http.StatusFound)
}

// HandleSteamCallback recebe o retorno da Steam após o login do usuário
// e executa o fluxo completo de autenticação:
//
//  1. Valida os parâmetros OpenID enviados pela Steam
//  2. Extrai o steamid64 do usuário
//  3. Busca nome, avatar e perfil na Steam Web API
//  4. Cria ou atualiza o usuário no MongoDB (upsert)
//  5. Gera um JWT interno com o ID do usuário no banco
//  6. Redireciona pro frontend com o token na query string
//
// Rota: GET /auth/steam/callback
func HandleSteamCallback(w http.ResponseWriter, r *http.Request) {
	baseURL := os.Getenv("APP_BASE_URL")
	frontendURL := os.Getenv("FRONTEND_URL")
	if baseURL == "" || frontendURL == "" {
		http.Error(w, "Configuração interna inválida", http.StatusInternalServerError)
		return
	}

	callbackURL := fmt.Sprintf("%s/auth/steam/callback", baseURL)

	// ── Passo 1: validar a resposta OpenID ──────────────────────────────────
	//
	// Verify valida a resposta da Steam.
	// Internamente ela:
	//   - Confere o nonce (evita replay attacks — impede que alguém reutilize
	//     uma URL de callback antiga pra se autenticar)
	//   - Faz uma requisição direta pra Steam confirmando os parâmetros
	//   - Retorna a identidade do usuário se tudo estiver ok
	identity, err := openid.Verify(
		callbackURL+"?"+r.URL.RawQuery,
		discoveryCache,
		nonceStore,
	)
	if err != nil {
		log.Printf("[steam] falha na verificação OpenID: %v", err)
		http.Error(w, "Falha na autenticação com a Steam", http.StatusUnauthorized)
		return
	}

	// ── Passo 2: extrair o steamid64 ────────────────────────────────────────
	//
	// A identidade tem o formato:
	// https://steamcommunity.com/openid/id/76561198XXXXXXXXX
	steamID := extractSteamID(identity)
	if steamID == "" {
		log.Printf("[steam] não foi possível extrair steamid de: %s", identity)
		http.Error(w, "SteamID inválido", http.StatusBadRequest)
		return
	}

	log.Printf("[steam] steamid validado: %s", steamID)

	// ── Passo 3: buscar perfil na Steam Web API ──────────────────────────────
	//
	// Com o steamid em mãos, chamamos a Steam API pra buscar
	// nome de exibição, avatar e URL do perfil do usuário.
	profile, err := service.GetSteamProfile(steamID)
	if err != nil {
		log.Printf("[steam] erro ao buscar perfil: %v", err)
		http.Error(w, "Erro ao buscar perfil na Steam", http.StatusInternalServerError)
		return
	}

	log.Printf("[steam] perfil obtido: %s (%s)", profile.Username, profile.SteamID)

	// ── Passo 4: criar ou atualizar usuário no MongoDB ───────────────────────
	//
	// UpsertSteamUser garante que:
	//   - Primeiro login → cria o documento no banco
	//   - Logins seguintes → atualiza nome/avatar (podem mudar na Steam)
	//     e registra o último login
	user, err := repository.UpsertSteamUser(model.User{
		SteamID:    profile.SteamID,
		Username:   profile.Username,
		AvatarURL:  profile.AvatarURL,
		ProfileURL: profile.ProfileURL,
	})
	if err != nil {
		log.Printf("[steam] erro ao salvar usuário: %v", err)
		http.Error(w, "Erro ao salvar usuário", http.StatusInternalServerError)
		return
	}

	// ── Passo 5: gerar JWT interno ───────────────────────────────────────────
	//
	// O JWT carrega o ID interno do usuário no MongoDB (não o steamid).
	// Isso desacopla nossa autenticação da Steam — se um dia adicionarmos
	// outro método de login (ex: GitHub), o resto do sistema não muda.
	token, err := util.GenerateJWT(user.ID.Hex())
	if err != nil {
		log.Printf("[steam] erro ao gerar JWT: %v", err)
		http.Error(w, "Erro ao gerar token", http.StatusInternalServerError)
		return
	}

	// ── Passo 6: redirecionar pro frontend com o token ───────────────────────
	//
	// Passamos o JWT como query param na URL do frontend.
	// O frontend vai ler esse token, salvar no localStorage (ou cookie)
	// e usar em todas as requisições autenticadas no header Authorization.
	redirectURL := fmt.Sprintf("%s/auth/callback?token=%s", frontendURL, token)
	http.Redirect(w, r, redirectURL, http.StatusFound)
}

// extractSteamID extrai o steamid64 da URL de identidade retornada pela Steam.
//
// Exemplo de entrada:
//
//	"https://steamcommunity.com/openid/id/76561198123456789"
//
// Exemplo de saída:
//
//	"76561198123456789"
func extractSteamID(identity string) string {
	const prefix = "https://steamcommunity.com/openid/id/"

	if !strings.HasPrefix(identity, prefix) {
		return ""
	}

	return strings.TrimPrefix(identity, prefix)
}
