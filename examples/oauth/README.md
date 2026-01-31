# OAuth2 Authorization Example

This example demonstrates how to obtain a Strava access token using OAuth2.

## Prerequisites

1. A Strava account
2. Create an application at https://www.strava.com/settings/api
3. Set the **Authorization Callback Domain** to `localhost`

## Setup

1. Get your credentials from https://www.strava.com/settings/api:
   - Client ID
   - Client Secret

2. Set environment variables:
```bash
export STRAVA_CLIENT_ID="your_client_id"
export STRAVA_CLIENT_SECRET="your_client_secret"
```

## Running

```bash
go run main.go
```

## What It Does

1. **Starts a local web server** on port 8080 to handle the OAuth callback
2. **Prints an authorization URL** - you need to open this in your browser
3. **Strava login page** - Log in and authorize the application
4. **Receives the authorization code** from Strava's redirect
5. **Exchanges the code** for an access token
6. **Prints the tokens** - save these for future use
7. **Tests the token** by fetching your activities

## Expected Output

```
=== Strava OAuth2 Authorization ===

Please open this URL in your browser:

https://www.strava.com/oauth/authorize?client_id=...

Waiting for authorization...

Exchanging authorization code for access token...

=== Success! ===

Access Token: your_access_token_here_abc123def456
Refresh Token: your_refresh_token_here_xyz789
Expires At: 2025-12-05 21:30:45

Athlete: John Doe

=== Add to your .env file or export: ===
export STRAVA_ACCESS_TOKEN="your_access_token_here"
export STRAVA_REFRESH_TOKEN="your_refresh_token_here"

=== Testing token by fetching activities ===

Found 10 activities. Latest:
- Morning Run (5.21 km)
```

## Token Management

### Access Token
- Valid for **6 hours**
- Use this to make API requests

### Refresh Token
- Does not expire
- Use this to get a new access token when it expires

## Refreshing an Expired Token

```go
package main

import (
    "context"
    "github.com/kpi-studio/go-strava-api/internal/auth"
)

func main() {
    oauth := &auth.OAuth2Config{
        ClientID:     "your_client_id",
        ClientSecret: "your_client_secret",
    }

    refreshToken := "your_refresh_token"

    // Get a new access token
    token, err := oauth.RefreshToken(context.Background(), refreshToken)
    if err != nil {
        panic(err)
    }

    // Use the new access token
    println("New Access Token:", token.AccessToken)
}
```

## Using TokenManager

For automatic token refresh:

```go
package main

import (
    "context"
    "github.com/kpi-studio/go-strava-api"
    "github.com/kpi-studio/go-strava-api/internal/auth"
)

func main() {
    oauth := &auth.OAuth2Config{
        ClientID:     "your_client_id",
        ClientSecret: "your_client_secret",
    }

    // Initial token from authorization
    initialToken := &auth.TokenResponse{
        AccessToken:  "current_access_token",
        RefreshToken: "refresh_token",
        ExpiresAt:    1234567890,
    }

    // Create token manager with auto-refresh
    tokenManager := auth.NewTokenManager(oauth, initialToken)

    // Set callback to save new tokens
    tokenManager.SetTokenUpdateCallback(func(token *auth.TokenResponse) {
        println("Token refreshed! Save this:")
        println("Access Token:", token.AccessToken)
        println("Refresh Token:", token.RefreshToken)
    })

    // Get access token (automatically refreshes if needed)
    ctx := context.Background()
    accessToken, err := tokenManager.GetAccessToken(ctx)
    if err != nil {
        panic(err)
    }

    // Use with client
    client := strava.NewClient(accessToken)
    activities, _ := client.Activities.List(ctx, nil)
    println("Activities:", len(activities))
}
```

## Scopes

Available scopes:
- `read` - Read public data
- `read_all` - Read private data
- `profile:read_all` - Read profile data
- `profile:write` - Update profile
- `activity:read` - Read activity data
- `activity:read_all` - Read all activity data
- `activity:write` - Create and update activities

Example with all scopes:
```go
oauth := &auth.OAuth2Config{
    ClientID:     clientID,
    ClientSecret: clientSecret,
    RedirectURI:  redirectURI,
    Scopes:       []string{"read", "activity:read_all", "activity:write"},
}
```
