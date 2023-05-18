package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetTwitterCredentials() (string, string, string, string) {
	//Pegando variáveis do ambiente
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	consumerKey := os.Getenv("consumer_key")
	consumerSecret := os.Getenv("consumer_secret")
	accessToken := os.Getenv("access_token")
	accessTokenSecret := os.Getenv("access_token_secret")

	return consumerKey, consumerSecret, accessToken, accessTokenSecret
}

func GetSpotifyCredentials() (string, string) {
	//Pegando variáveis do ambiente
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	clientID := os.Getenv("client_id")
	clientSecret := os.Getenv("client_secret")

	return clientID, clientSecret
}

func GetSpotifyRefreshToken() string {
	//Pegando variáveis do ambiente
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	refreshToken := os.Getenv("refresh_token")

	return refreshToken
}
