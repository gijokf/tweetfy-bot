package controller

import (
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"time"

	spotify "twitter-bot-1.0/services/Spotify"
)

type TrackItem struct {
	Track struct {
		Name         string `json:"name"`
		ExternalURLs struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
	} `json:"track"`
}

type TrackResponse struct {
	Items []TrackItem `json:"items"`
}

func GetSavedTrack() string {
	accessToken := spotify.RenewToken()

	url := "https://api.spotify.com/v1/me/player/recently-played"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", "Bearer "+accessToken)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var trackResponse TrackResponse
	err = json.Unmarshal(body, &trackResponse)
	if err != nil {
		panic(err)
	}

	if len(trackResponse.Items) > 0 {
		rand.NewSource(time.Now().Unix())
		randomIndex := rand.Intn(len(trackResponse.Items))

		trackItem := trackResponse.Items[randomIndex]

		return trackItem.Track.ExternalURLs.Spotify
	} else {
		return "No tracks found."
	}
}
