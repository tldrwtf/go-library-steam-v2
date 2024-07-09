package steam

import (
	"encoding/json"
	"fmt"
)

// Item represents an item in a trade or market listing
type Item struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// SendTradeOffer sends a trade offer to another user
func (bot *Bot) SendTradeOffer(steamID string, items []Item) error {
	itemData, err := json.Marshal(items)
	if err != nil {
		return fmt.Errorf("failed to marshal items: %w", err)
	}
	data := map[string]string{
		"key":     bot.APIKey,
		"steamid": steamID,
		"items":   string(itemData),
	}
	_, err = bot.makeAPIPostRequest("ISteamTrading/SendTradeOffer/v1", data)
	if err != nil {
		return fmt.Errorf("failed to send trade offer: %w", err)
	}
	return nil
}

// ListItemsOnMarket lists items on the Steam market
func (bot *Bot) ListItemsOnMarket(items []Item) error {
	itemData, err := json.Marshal(items)
	if err != nil {
		return fmt.Errorf("failed to marshal items: %w", err)
	}
	data := map[string]string{
		"key":   bot.APIKey,
		"items": string(itemData),
	}
	_, err = bot.makeAPIPostRequest("ISteamEconomy/ListItemsOnMarket/v1", data)
	if err != nil {
		return fmt.Errorf("failed to list items on market: %w", err)
	}
	return nil
}
