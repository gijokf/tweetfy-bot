package spotify

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"twitter-bot-1.0/config"
	responseBody "twitter-bot-1.0/models"
)

const (
	redirectURI  = "http://localhost:8888/callback"
	authorizeURL = "https://accounts.spotify.com/authorize"
)

var (
	clientID, clientSecret = config.GetSpotifyCredentials()
)

func Server() {
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/callback", CallbackHandler)
	http.ListenAndServe(":8888", nil)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	state, err := GenerateRandomString(16)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	scope := "user-read-private user-library-read user-read-recently-played"
	values := url.Values{}
	values.Set("response_type", "code")
	values.Set("client_id", clientID)
	values.Set("redirect_uri", redirectURI)
	values.Set("state", state)
	values.Set("scope", scope)
	http.Redirect(w, r, authorizeURL+"?"+values.Encode(), http.StatusFound)
}

func CallbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")
	state := r.FormValue("state")

	GenerateToken(code)
	fmt.Fprintf(w, "Authorization code: %s\nState: %s", code, state)
}

func GenerateRandomString(n int) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func GenerateToken(code string) {
	// Montar body
	bodyRequest := url.Values{
		"grant_type":   {"authorization_code"},
		"code":         {code},
		"redirect_uri": {"http://localhost:8888/callback"},
	}

	// Cria o request
	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(bodyRequest.Encode()))
	if err != nil {
		panic(err)
	}

	auth := clientID + ":" + clientSecret
	authBase64 := base64.StdEncoding.EncodeToString([]byte(auth))

	// Define o header com a autenticação
	req.Header.Set("Authorization", "Basic "+authBase64)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Faz a requisição
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Lê o corpo da resposta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Decodifica o JSON da resposta
	var responseBody responseBody.ResponseBody
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		panic(err)
	}

	fileAccess, _ := os.Create("access_token.txt")
	fmt.Fprintln(fileAccess, responseBody.AccessToken)

	fileRefresh, _ := os.Create("refresh_token.txt")
	fmt.Fprintln(fileRefresh, responseBody.RefreshToken)
}

func RenewToken() string {
	refreshToken := config.GetSpotifyRefreshToken()

	// Montar body
	bodyRequest := url.Values{
		"grant_type":    {"refresh_token"},
		"refresh_token": {refreshToken},
	}

	// Cria o request
	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(bodyRequest.Encode()))
	if err != nil {
		panic(err)
	}

	auth := clientID + ":" + clientSecret
	authBase64 := base64.StdEncoding.EncodeToString([]byte(auth))

	// Define o header com a autenticação
	req.Header.Set("Authorization", "Basic "+authBase64)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Faz a requisição
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Lê o corpo da resposta
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Decodifica o JSON da resposta
	var responseBody responseBody.ResponseBody
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		panic(err)
	}

	return responseBody.AccessToken
}
