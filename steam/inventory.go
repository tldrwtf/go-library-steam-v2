package steam

import (
	"encoding/json"
	"fmt"
)

// Inventory represents a user's inventory
type Inventory struct {
	Items []Item `json:"items"`
}

// GetInventory retrieves the inventory for a given Steam ID
func (bot *Bot) GetInventory(steamID string) (*Inventory, error) {
	params := map[string]string{
		"key":     bot.APIKey,
		"steamid": steamID,
	}
	body, err := bot.makeAPIRequest("ISteamUser/GetInventory/v1", params)
	if err != nil {
		return nil, fmt.Errorf("failed to get inventory: %w", err)
	}

	var inventory Inventory
	err = json.Unmarshal(body, &inventory)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal inventory: %w", err)
	}
	return &inventory, nil
}
