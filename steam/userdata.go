package steam

import (
	"encoding/json"
	"fmt"
)

// UserData represents a user's profile data
type UserData struct {
	SteamID      string `json:"steamid"`
	PersonaName  string `json:"personaname"`
	ProfileURL   string `json:"profileurl"`
	Avatar       string `json:"avatar"`
	AvatarMedium string `json:"avatarmedium"`
	AvatarFull   string `json:"avatarfull"`
}

// GetUserData retrieves the profile data for a given Steam ID
func (bot *Bot) GetUserData(steamID string) (*UserData, error) {
	params := map[string]string{
		"key":      bot.APIKey,
		"steamids": steamID,
	}
	body, err := bot.makeAPIRequest("ISteamUser/GetPlayerSummaries/v2", params)
	if err != nil {
		return nil, fmt.Errorf("failed to get user data: %w", err)
	}

	var result struct {
		Response struct {
			Players []UserData `json:"players"`
		} `json:"response"`
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal user data: %w", err)
	}
	if len(result.Response.Players) == 0 {
		return nil, fmt.Errorf("no user data found for steamID: %s", steamID)
	}
	return &result.Response.Players[0], nil
}
