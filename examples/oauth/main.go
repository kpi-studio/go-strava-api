package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/kpi-studio/go-strava-api"
	"github.com/kpi-studio/go-strava-api/internal/auth"
)

func main() {
	// Get credentials from environment variables
	clientID := os.Getenv("STRAVA_CLIENT_ID")
	clientSecret := os.Getenv("STRAVA_CLIENT_SECRET")
	redirectURI := "http://localhost:8080/callback"

	if clientID == "" || clientSecret == "" {
		log.Fatal("STRAVA_CLIENT_ID and STRAVA_CLIENT_SECRET environment variables are required")
	}

	// Create OAuth2 config
	oauth := &auth.OAuth2Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURI:  redirectURI,
		Scopes:       []string{"read", "activity:read_all"},
	}

	// Channel to receive the authorization code
	codeChan := make(chan string)

	// Setup HTTP server to handle the callback
	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		if code == "" {
			http.Error(w, "No code in callback", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, `
			<html>
			<body>
				<h1>Authorization Successful!</h1>
				<p>You can close this window and return to your terminal.</p>
			</body>
			</html>
		`)

		codeChan <- code
	})

	// Start server in background
	server := &http.Server{Addr: ":8080"}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Generate authorization URL
	authURL := oauth.GetAuthorizationURL(auth.AuthorizationURLParams{
		ApprovalPrompt: "auto",
	})

	fmt.Println("\n=== Strava OAuth2 Authorization ===")
	fmt.Println("\nPlease open this URL in your browser:")
	fmt.Println("\n" + authURL + "\n")
	fmt.Println("Waiting for authorization...")

	// Wait for the authorization code
	code := <-codeChan

	// Shutdown the server
	ctx := context.Background()
	server.Shutdown(ctx)

	fmt.Println("\nExchanging authorization code for access token...")

	// Exchange the code for an access token
	token, err := oauth.ExchangeCode(ctx, code)
	if err != nil {
		log.Fatalf("Failed to exchange code: %v", err)
	}

	fmt.Println("\n=== Success! ===")
	fmt.Printf("\nAccess Token: %s\n", token.AccessToken)
	fmt.Printf("Refresh Token: %s\n", token.RefreshToken)
	fmt.Printf("Expires At: %s\n", token.ExpirationTime())
	fmt.Printf("\nAthlete: %s %s\n", token.Athlete.Firstname, token.Athlete.Lastname)

	// Save to environment variable format
	fmt.Println("\n=== Add to your .env file or export: ===")
	fmt.Printf("export STRAVA_ACCESS_TOKEN=\"%s\"\n", token.AccessToken)
	fmt.Printf("export STRAVA_REFRESH_TOKEN=\"%s\"\n", token.RefreshToken)

	// Test the token by fetching activities
	fmt.Println("\n=== Testing token by fetching activities ===")
	client := strava.NewClient(token.AccessToken)
	activities, err := client.Activities.List(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to fetch activities: %v", err)
	}

	fmt.Printf("\nFound %d activities. Latest:\n", len(activities))
	if len(activities) > 0 {
		activity := activities[0]
		fmt.Printf("- %s (%.2f km)\n", activity.Name, activity.Distance/1000)
	}
}
