package main

import (
	"twitter-bot-1.0/controller"
	spotify "twitter-bot-1.0/services/Spotify"
)

func main() {
	spotify.Server()
	controller.CreateTweet()
}
