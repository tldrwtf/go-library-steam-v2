package main

import (
	"go-library-steam/config"
	"go-library-steam/steam"
	"log"
)

func main() {
	log.Println("Starting application...")

	err := config.LoadEnv(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	err = config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	apiKey := config.GetEnv("STEAM_API_KEY")
	steamGuard := &steam.SteamGuard{
		SharedSecret:   config.GetEnv("STEAM_SHARED_SECRET"),
		IdentitySecret: config.GetEnv("STEAM_IDENTITY_SECRET"),
	}

	bot := steam.NewBot(apiKey, steamGuard)

	username := config.GetEnv("STEAM_USERNAME")
	password := config.GetEnv("STEAM_PASSWORD")

	err = bot.Login(username, password)
	if err != nil {
		log.Fatalf("Error logging in: %v", err)
	}

	/* Example usage
	err = bot.AddFriend("someSteamID")
	if err != nil {
		log.Fatalf("Error adding friend: %v", err)
	}
	log.Println("Friend added successfully") */
}
