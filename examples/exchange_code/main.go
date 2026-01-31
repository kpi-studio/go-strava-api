package main

import (
	"context"
	"fmt"
	"log"

	"github.com/kpi-studio/strava-api/internal/auth"
)

func main() {
	// FROM: https://www.strava.com/settings/api
	clientID := "YOUR_CLIENT_ID"         // Replace with your Client ID
	clientSecret := "YOUR_CLIENT_SECRET" // Replace with your Client Secret

	// The authorization code from the callback URL
	authorizationCode := "YOUR_AUTHORIZATION_CODE"

	if clientID == "YOUR_CLIENT_ID" || clientSecret == "YOUR_CLIENT_SECRET" {
		log.Fatal(`
Please edit this file and replace:
- YOUR_CLIENT_ID with your actual Client ID
- YOUR_CLIENT_SECRET with your actual Client Secret

Get these from: https://www.strava.com/settings/api
`)
	}

	// Create OAuth config
	oauth := &auth.OAuth2Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURI:  "http://localhost:8080/callback",
	}

	fmt.Println("Exchanging authorization code for access token...")
	fmt.Println("Code:", authorizationCode)
	fmt.Println()

	// Exchange the code for an access token
	ctx := context.Background()
	token, err := oauth.ExchangeCode(ctx, authorizationCode)
	if err != nil {
		log.Fatalf("Failed to exchange code: %v\n\nMake sure:\n1. Your Client ID and Secret are correct\n2. The authorization code hasn't been used already (it's single-use)\n3. The code hasn't expired (expires in 10 minutes)\n", err)
	}

	fmt.Println("✅ SUCCESS!")
	fmt.Println()
	fmt.Println("=== YOUR ACCESS TOKEN ===")
	fmt.Println(token.AccessToken)
	fmt.Println()
	fmt.Println("=== YOUR REFRESH TOKEN ===")
	fmt.Println(token.RefreshToken)
	fmt.Println()
	fmt.Println("=== Athlete Info ===")
	fmt.Printf("Name: %s %s\n", token.Athlete.Firstname, token.Athlete.Lastname)
	fmt.Printf("ID: %d\n", token.Athlete.ID)
	fmt.Printf("Token expires at: %s\n", token.ExpirationTime())
	fmt.Println()
	fmt.Println("=== Copy this to your terminal ===")
	fmt.Printf("export STRAVA_ACCESS_TOKEN=\"%s\"\n", token.AccessToken)
	fmt.Println()
	fmt.Println("Now you can run:")
	fmt.Println("cd ../get_activities && go run main.go")
}
