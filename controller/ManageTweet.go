package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"tweetfy-bot/config"

	"github.com/dghubble/oauth1"
)

var (
	consumerKey, consumerSecret, accessToken, accessTokenSecret = config.GetTwitterCredentials()
	spotifyLink                                                 = GetSavedTrack()
)

func CreateTweet(text string) {
	// Configurar as credenciais de autenticação
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessTokenSecret)
	httpClient := config.Client(context.Background(), token)

	// Configurar o corpo da requisição
	data := make(map[string]string)

	data["text"] = text + "\n" + spotifyLink

	// hour := time.Now().Hour()

	// switch {
	// case hour < 12:
	// 	data["text"] = "Bom dia do bot 👋🏻🤖.\n" + spotifyLink
	// case hour < 18:
	// 	data["text"] = "Boa tarde do bot 👋🏻🤖.\n" + spotifyLink
	// default:
	// 	data["text"] = "Boa noite do bot 👋🏻🤖.\n" + spotifyLink
	// }

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Erro ao codificar o corpo da requisição em JSON:", err)
		os.Exit(1)
	}

	// Configurar a requisição HTTP
	req, err := http.NewRequest("POST", "https://api.twitter.com/2/tweets", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Erro ao criar a requisição HTTP:", err)
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/json")

	// Enviar a requisição HTTP
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println("Erro ao enviar a requisição HTTP:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Ler a resposta da requisição HTTP
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler a resposta da requisição HTTP:", err)
		os.Exit(1)
	}

	fmt.Println(string(respBody))

	// Exibir o ID do tweet criado
	var tweetData map[string]interface{}
	err = json.Unmarshal(respBody, &tweetData)
	if err != nil {
		fmt.Println("Erro ao decodificar a resposta da requisição HTTP:", err)
		os.Exit(1)
	}

	tweetID := tweetData["data"].(map[string]interface{})["id"].(string)
	fmt.Println("Tweet criado com sucesso. ID:", tweetID)
}
