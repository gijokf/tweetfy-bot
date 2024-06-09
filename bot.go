package main

import (
	"tweetfy-bot/controller"
	spotify "tweetfy-bot/services/Spotify"
)

func main() {
	spotify.Server()
	controller.CreateTweet()
}
