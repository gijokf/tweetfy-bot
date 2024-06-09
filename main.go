package main

import (
	"fmt"
	"tweetfy-bot/controller"
	"tweetfy-bot/services"
	spotify "tweetfy-bot/services/Spotify"
)

func main() {
	spotify.Server()

	text, err := services.TweetFile()
	if err != nil {
		fmt.Println("Erro na leitura do arquivo: ", err)
	}

	controller.CreateTweet(text)
}
