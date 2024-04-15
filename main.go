package main

import (
	"context"
	"encoding/json"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	apiClientID := os.Getenv("BNET_API_CLIENT_ID")
	apiClientSecret := os.Getenv("BNET_API_CLIENT_SECRET")

	ctx := context.TODO()

	clCfg := &clientcredentials.Config{
		ClientID:     apiClientID,
		ClientSecret: apiClientSecret,
		TokenURL:     "https://oauth.battle.net/token",
	}
	httpClient := clCfg.Client(ctx)

	var (
		req  *http.Request
		res  *http.Response
		body []byte
		err  error
	)

	req, err = http.NewRequestWithContext(ctx, "POST", clCfg.TokenURL, strings.NewReader("grant_type=client_credentials"))
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	q.Set("grant_type", "client_credentials")
	req.URL.RawQuery = q.Encode()

	req.SetBasicAuth(apiClientID, apiClientSecret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err = httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Error closing response body: %v", err)
		}
	}(res.Body)

	body, err = io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var token *oauth2.Token
	err = json.Unmarshal(body, &token)
	if err != nil {
		log.Fatal(err)
	}

}
