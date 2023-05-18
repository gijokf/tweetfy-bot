package models

type Track struct {
	ExternalURLs struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
}
