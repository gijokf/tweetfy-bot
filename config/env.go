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

	twitterConsumerKey := os.Getenv("twitter_consumer_key")
	twitterConsumerSecret := os.Getenv("twitter_consumer_secret")
	twitterAccessToken := os.Getenv("twitter_access_token")
	twitterAccessTokenSecret := os.Getenv("twitter_access_token_secret")

	return twitterConsumerKey, twitterConsumerSecret, twitterAccessToken, twitterAccessTokenSecret
}

func GetSpotifyCredentials() (string, string) {
	//Pegando variáveis do ambiente
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	spotifyClientID := os.Getenv("spotify_client_id")
	spotifyClientSecret := os.Getenv("spotify_client_secret")

	return spotifyClientID, spotifyClientSecret
}

func GetSpotifyRefreshToken() string {
	//Pegando variáveis do ambiente
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	spotifyRefreshToken := os.Getenv("spotify_refresh_token")

	return spotifyRefreshToken
}
