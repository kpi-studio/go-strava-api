# Getting Started with Strava API

## How to Get an Access Token

There are two ways to get a Strava access token:

### Method 1: Quick Testing (Temporary Token)

**Fastest way for testing - expires in 6 hours**

1. Go to https://www.strava.com/settings/api
2. Create an application (if you haven't already)
3. Copy the **Access Token** shown on that page
4. Use it directly:

```bash
export STRAVA_ACCESS_TOKEN="your_token_here"
cd examples/get_activities
go run main.go
```

### Method 2: OAuth2 Flow (Production)

**Recommended for production - includes refresh token**

#### Setup:

1. **Create a Strava Application**
   - Go to https://www.strava.com/settings/api
   - Click "Create an App"
   - Fill in the details:
     - **Authorization Callback Domain**: `localhost`
   - Note your **Client ID** and **Client Secret**

2. **Set environment variables:**
   ```bash
   export STRAVA_CLIENT_ID="your_client_id"
   export STRAVA_CLIENT_SECRET="your_client_secret"
   ```

3. **Run the OAuth example:**
   ```bash
   cd examples/oauth
   go run main.go
   ```

4. **Follow the instructions:**
   - Open the URL shown in your browser
   - Log in to Strava
   - Authorize the application
   - The token will be displayed in your terminal

5. **Save the tokens:**
   ```bash
   export STRAVA_ACCESS_TOKEN="the_access_token_from_output"
   export STRAVA_REFRESH_TOKEN="the_refresh_token_from_output"
   ```

## Running Examples

### Get Activities

Fetches your recent activities:

```bash
cd examples/get_activities
export STRAVA_ACCESS_TOKEN="your_token"
go run main.go
```

**Output:**
```
Fetching activities...

Found 10 activities:

1. Morning Run
   Type: Run
   Distance: 5.21 km
   Moving Time: 28 minutes
   Date: 2025-12-04 07:30

...
```

## Using the Library

### Basic Usage

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/kpi-studio/strava-api"
    "github.com/kpi-studio/strava-api/models"
)

func main() {
    // Create client
    client := strava.NewClient("your_access_token")
    ctx := context.Background()

    // Get activities
    activities, err := client.Activities.List(ctx, &models.ListOptions{
        Page:    1,
        PerPage: 10,
    })
    if err != nil {
        log.Fatal(err)
    }

    for _, activity := range activities {
        fmt.Printf("%s - %.2f km\n", activity.Name, activity.Distance/1000)
    }
}
```

### With Automatic Token Refresh

```go
package main

import (
    "context"
    "github.com/kpi-studio/strava-api"
    "github.com/kpi-studio/strava-api/internal/auth"
)

func main() {
    // OAuth config
    oauth := &auth.OAuth2Config{
        ClientID:     "your_client_id",
        ClientSecret: "your_client_secret",
    }

    // Initial token
    token := &auth.TokenResponse{
        AccessToken:  "current_access_token",
        RefreshToken: "refresh_token",
        ExpiresAt:    1234567890,
    }

    // Create token manager
    tokenManager := auth.NewTokenManager(oauth, token)

    // Optional: Save new tokens when refreshed
    tokenManager.SetTokenUpdateCallback(func(newToken *auth.TokenResponse) {
        // Save to database or file
        fmt.Println("Token refreshed:", newToken.AccessToken)
    })

    // Get valid access token (auto-refreshes if expired)
    ctx := context.Background()
    accessToken, _ := tokenManager.GetAccessToken(ctx)

    // Use with client
    client := strava.NewClient(accessToken)
    activities, _ := client.Activities.List(ctx, nil)
    fmt.Println("Found", len(activities), "activities")
}
```

## Available Services

The client provides access to these services:

```go
client := strava.NewClient("token")

// Activities
client.Activities.List(ctx, opts)
client.Activities.Get(ctx, activityID, includeEfforts)
client.Activities.Create(ctx, params)
client.Activities.Update(ctx, activityID, update)
client.Activities.Delete(ctx, activityID)

// Athletes
client.Athletes.GetCurrent(ctx)
client.Athletes.Get(ctx, athleteID)
client.Athletes.GetStats(ctx, athleteID)
client.Athletes.UpdateWeight(ctx, weight)

// Clubs
client.Clubs.Get(ctx, clubID)
client.Clubs.ListMembers(ctx, clubID, pagination)
client.Clubs.ListActivities(ctx, clubID, opts)

// Segments
client.Segments.Get(ctx, segmentID)
client.Segments.ListEfforts(ctx, segmentID, opts)
client.Segments.GetLeaderboard(ctx, segmentID, opts)
client.Segments.Explore(ctx, opts)

// And more: Routes, Gears, Streams, Uploads
```

## Token Scopes

When requesting authorization, specify the scopes you need:

- `read` - Read public data
- `read_all` - Read private data
- `activity:read` - Read activity data
- `activity:read_all` - Read all activities (including private)
- `activity:write` - Create and update activities
- `profile:read_all` - Read all profile information
- `profile:write` - Update profile

Example:
```go
oauth := &auth.OAuth2Config{
    ClientID:     clientID,
    ClientSecret: clientSecret,
    RedirectURI:  redirectURI,
    Scopes:       []string{"read", "activity:read_all", "activity:write"},
}
```

## Troubleshooting

### "Invalid token" error
- Check if your token expired (tokens expire after 6 hours)
- Use the refresh token to get a new access token
- Or re-run the OAuth flow

### "Rate limit exceeded"
- Strava has rate limits: 100 requests per 15 minutes, 1000 per day
- The library includes automatic rate limiting
- Wait before making more requests

### "Authorization callback domain" error
- Make sure you set `localhost` as the callback domain in your Strava app settings
- The redirect URI must be `http://localhost:8080/callback`

## Next Steps

- Check out the [examples](./examples/) directory
- Read the [API documentation](https://developers.strava.com/docs/reference/)
