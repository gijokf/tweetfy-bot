package models

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
