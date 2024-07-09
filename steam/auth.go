package steam

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// Login performs the login action using the provided username and password
func (bot *Bot) Login(username, password string) error {
	url := "https://steamcommunity.com/login/dologin/"

	data := map[string]string{
		"username": username,
		"password": password,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return errors.New("failed to login")
	}

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return err
	}

	if !result["success"].(bool) {
		return errors.New("login failed")
	}

	return nil
}
