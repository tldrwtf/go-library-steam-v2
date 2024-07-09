package steam

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Bot struct {
	APIKey     string
	SteamGuard *Guard
}

// makeAPIRequest sends a GET request to the Steam API
func (bot *Bot) makeAPIRequest(endpoint string, params map[string]string) ([]byte, error) {
	url := fmt.Sprintf("https://api.steampowered.com/%s", endpoint)
	if len(params) > 0 {
		paramStr := "?"
		for key, value := range params {
			paramStr += fmt.Sprintf("%s=%s&", key, value)
		}
		url += paramStr[:len(paramStr)-1]
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make API request: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read API response: %w", err)
	}

	return body, nil
}

// makeAPIPostRequest sends a POST request to the Steam API
func (bot *Bot) makeAPIPostRequest(endpoint string, data map[string]string) ([]byte, error) {
	url := fmt.Sprintf("https://api.steampowered.com/%s", endpoint)
	formData := ""
	for key, value := range data {
		formData += fmt.Sprintf("%s=%s&", key, value)
	}
	formData = strings.TrimSuffix(formData, "&")

	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(formData))
	if err != nil {
		return nil, fmt.Errorf("failed to make API POST request: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API POST request failed with status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read API response: %w", err)
	}

	return body, nil
}
