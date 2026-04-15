package model

// SteamGame representa um jogo retornado pela Steam Web API
// no endpoint GetOwnedGames.
//
// Esses dados são específicos da Steam e não se misturam com o
// model Game genérico do sistema (que cobre jogos adicionados manualmente).
// Na hora de importar a biblioteca, SteamGame é convertido para LibraryItem.
type SteamGame struct {
	// AppID é o identificador único do jogo na Steam.
	AppID uint32 `json:"appid"`

	// Name é o nome do jogo — só retornado se include_appinfo=true.
	Name string `json:"name"`

	// PlaytimeForever é o total de minutos jogados desde que a Steam
	// começou a registrar (início de 2009).
	PlaytimeForever int `json:"playtime_forever"`

	// Playtime2Weeks é o total de minutos jogados nas últimas 2 semanas.
	// Pode ser 0 se o jogo não foi jogado recentemente.
	Playtime2Weeks int `json:"playtime_2weeks"`

	// ImgIconURL é o hash do ícone do jogo.
	// URL completa: https://media.steampowered.com/steamcommunity/public/images/apps/{appid}/{hash}.jpg
	ImgIconURL string `json:"img_icon_url"`

	// ImgLogoURL é o hash do logo do jogo.
	// URL completa: https://media.steampowered.com/steamcommunity/public/images/apps/{appid}/{hash}.jpg
	ImgLogoURL string `json:"img_logo_url"`

	// HasCommunityVisibleStats indica se existe uma página de stats/achievements
	// para este jogo em: https://steamcommunity.com/profiles/{steamid}/stats/{appid}
	HasCommunityVisibleStats bool `json:"has_community_visible_stats"`
}

// IconURL monta a URL completa do ícone do jogo a partir do hash retornado pela Steam.
func (g SteamGame) IconURL() string {
	if g.ImgIconURL == "" {
		return ""
	}
	return formatSteamImageURL(g.AppID, g.ImgIconURL)
}

// LogoURL monta a URL completa do logo do jogo a partir do hash retornado pela Steam.
func (g SteamGame) LogoURL() string {
	if g.ImgLogoURL == "" {
		return ""
	}
	return formatSteamImageURL(g.AppID, g.ImgLogoURL)
}

func formatSteamImageURL(appID uint32, hash string) string {
	return "https://media.steampowered.com/steamcommunity/public/images/apps/" +
		itoa(appID) + "/" + hash + ".jpg"
}

// itoa converte uint32 pra string sem importar strconv no model.
func itoa(n uint32) string {
	if n == 0 {
		return "0"
	}
	buf := [10]byte{}
	pos := len(buf)
	for n > 0 {
		pos--
		buf[pos] = byte(n%10) + '0'
		n /= 10
	}
	return string(buf[pos:])
}
