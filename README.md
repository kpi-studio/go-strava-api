# Strava API Go Client

A comprehensive, easy-to-use Go client library for the Strava API v3.

## Features

- Complete coverage of Strava API v3 endpoints
- OAuth 2.0 authentication with automatic token refresh
- Built-in rate limiting and retry logic
- Strong type safety with comprehensive structs
- Context support for cancellation and timeouts
- Utility functions for common conversions
- Polyline encoding/decoding
- Comprehensive examples

## Installation

```bash
go get github.com/kpi-studio/go-strava-api
```

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"

    strava "github.com/kpi-studio/go-strava-api"
)

func main() {
    // Create a client with your access token
    client := strava.NewClient("YOUR_ACCESS_TOKEN")

    // Get the authenticated athlete
    ctx := context.Background()
    athlete, err := client.Athletes.GetCurrent(ctx)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Hello %s %s!\n", athlete.Firstname, athlete.Lastname)

    // List recent activities
    activities, err := client.Activities.List(ctx, &strava.ListOptions{
        Page:    1,
        PerPage: 10,
    })
    if err != nil {
        log.Fatal(err)
    }

    for _, activity := range activities {
        fmt.Printf("%s: %.2f km\n",
            activity.Name,
            strava.MetersToKilometers(activity.Distance))
    }
}
```

## Authentication

### OAuth 2.0 Flow

```go
// Create OAuth configuration
oauthConfig := &strava.OAuth2Config{
    ClientID:     "YOUR_CLIENT_ID",
    ClientSecret: "YOUR_CLIENT_SECRET",
    RedirectURI:  "http://localhost:8080/callback",
    Scopes: []string{
        "read",
        "activity:read_all",
        "activity:write",
    },
}

// Generate authorization URL
authURL := oauthConfig.GetAuthorizationURL(strava.AuthorizationURLParams{
    ApprovalPrompt: "auto",
    State:          "mystate",
})

// After user authorizes, exchange code for token
token, err := oauthConfig.ExchangeCode(ctx, authorizationCode)
if err != nil {
    log.Fatal(err)
}

// Create client with access token
client := strava.NewClient(token.AccessToken)
```

### Automatic Token Refresh

```go
// Create token manager for automatic refresh
tokenManager := strava.NewTokenManager(oauthConfig, token)

// Set callback for token updates (e.g., to save to database)
tokenManager.SetTokenUpdateCallback(func(newToken *strava.TokenResponse) {
    saveTokenToDatabase(newToken)
})

// Get valid access token (refreshes automatically if needed)
accessToken, err := tokenManager.GetAccessToken(ctx)
```

## API Services

### Activities

```go
// List activities
activities, err := client.Activities.List(ctx, &strava.ListOptions{
    Before:  1609459200, // Unix timestamp
    After:   1577836800, // Unix timestamp
    Page:    1,
    PerPage: 30,
})

// Get detailed activity
activity, err := client.Activities.Get(ctx, activityID, true)

// Create manual activity
newActivity, err := client.Activities.Create(ctx, strava.CreateActivityParams{
    Name:           "Morning Run",
    Type:           strava.ActivityTypeRun,
    StartDateLocal: time.Now(),
    ElapsedTime:    1800, // seconds
    Distance:       5000, // meters
})

// Update activity
updated, err := client.Activities.Update(ctx, activityID, &strava.UpdatableActivity{
    Name:        "New Name",
    Description: "Updated description",
})

// Delete activity
err := client.Activities.Delete(ctx, activityID)

// Activity streams
streams, err := client.Streams.GetActivityStreams(ctx, activityID,
    []strava.StreamType{
        strava.StreamTypeTime,
        strava.StreamTypeDistance,
        strava.StreamTypeAltitude,
        strava.StreamTypeHeartrate,
    },
    "medium", // resolution
)
```

### Athletes

```go
// Get authenticated athlete
athlete, err := client.Athletes.GetCurrent(ctx)

// Get another athlete
otherAthlete, err := client.Athletes.Get(ctx, athleteID)

// Get athlete stats
stats, err := client.Athletes.GetStats(ctx, athleteID)

// Update athlete weight
updated, err := client.Athletes.UpdateWeight(ctx, 70.5) // kg
```

### Segments

```go
// Get segment
segment, err := client.Segments.Get(ctx, segmentID)

// Star/unstar segment
starred, err := client.Segments.Star(ctx, segmentID, true)

// Get segment leaderboard
leaderboard, err := client.Segments.GetLeaderboard(ctx, segmentID,
    &strava.LeaderboardOptions{
        Gender:    "M",
        AgeGroup:  "25_34",
        Following: true,
        Page:      1,
        PerPage:   10,
    })

// Explore segments
results, err := client.Segments.Explore(ctx, strava.ExploreOptions{
    Bounds:       []float64{37.821362, -122.505373, 37.842038, -122.465977},
    ActivityType: "running",
})
```

### Clubs

```go
// Get club
club, err := client.Clubs.Get(ctx, clubID)

// List club members
members, err := client.Clubs.ListMembers(ctx, clubID, &strava.Pagination{
    Page:    1,
    PerPage: 50,
})

// Join/leave club
membership, err := client.Clubs.Join(ctx, clubID)
membership, err := client.Clubs.Leave(ctx, clubID)
```

### Routes

```go
// Get route
route, err := client.Routes.Get(ctx, routeID)

// Export route as GPX
gpxData, err := client.Routes.GetGPX(ctx, routeID)

// List athlete's routes
routes, err := client.Routes.ListByAthlete(ctx, athleteID, nil)
```

### Uploads

```go
// Upload activity file
upload, err := client.Uploads.Upload(ctx, &strava.UploadOptions{
    File:         fileReader,
    DataType:     "fit",
    Name:         "My Activity",
    Description:  "Great ride!",
    ActivityType: strava.ActivityTypeRide,
})

// Check upload status
status, err := client.Uploads.GetUploadStatus(ctx, uploadID)

// Wait for upload completion
result, err := client.Uploads.WaitForUpload(ctx, uploadID, 5*time.Minute)
```

## Utility Functions

### Distance Conversions

```go
miles := strava.MetersToMiles(meters)
km := strava.MetersToKilometers(meters)
meters := strava.MilesToMeters(miles)
meters := strava.KilometersToMeters(km)
```

### Speed Conversions

```go
mph := strava.MetersPerSecondToMilesPerHour(mps)
kmh := strava.MetersPerSecondToKilometersPerHour(mps)
```

### Pace Calculations

```go
pacePerMile := strava.CalculatePacePerMile(distanceMeters, timeSeconds)
pacePerKm := strava.CalculatePacePerKilometer(distanceMeters, timeSeconds)
formatted := strava.FormatPace(paceSeconds) // "6:30"
```

### Time Formatting

```go
formatted := strava.FormatDuration(seconds) // "1:23:45" or "45:30"
seconds := strava.ParseDuration("1:23:45")
```

### Polyline Encoding/Decoding

```go
// Decode polyline to coordinates
coords := strava.DecodePolyline(encodedPolyline)

// Encode coordinates to polyline
polyline := strava.EncodePolyline(coordinates)
```

### Distance Calculation

```go
// Calculate distance between two points (Haversine formula)
distance := strava.CalculateDistance(lat1, lng1, lat2, lng2)
```

## Rate Limiting

The client includes automatic rate limiting with configurable options:

```go
client := strava.NewClientWithOptions(accessToken, strava.ClientOptions{
    RateLimit: &strava.RateLimiterConfig{
        Enabled:    true,
        MinDelay:   100 * time.Millisecond,
        MaxRetries: 3,
    },
})
```

## Error Handling

```go
// Check for specific error types
if err != nil {
    if strava.IsRateLimitError(err) {
        // Handle rate limit
    } else if strava.IsAuthError(err) {
        // Handle authentication error
    } else if strava.IsNotFoundError(err) {
        // Handle not found
    }
}

// Get detailed error information
if apiErr, ok := err.(*strava.Error); ok {
    fmt.Printf("API Error: %s (Status: %d)\n", apiErr.Message, apiErr.StatusCode)
}
```

## Advanced Configuration

```go
// Create client with custom HTTP client and base URL
client := strava.NewClientWithOptions(accessToken, strava.ClientOptions{
    HTTPClient: &http.Client{
        Timeout: 60 * time.Second,
    },
    BaseURL: "https://custom-api-endpoint.com",
    RateLimit: &strava.RateLimiterConfig{
        Enabled:    true,
        MinDelay:   200 * time.Millisecond,
        MaxRetries: 5,
    },
})
```

## Examples

See the [`examples`](examples/) directory for complete working examples:

- [`basic`](examples/basic/) - Basic usage and athlete information
- [`oauth`](examples/oauth/) - OAuth authentication flow
- [`activities`](examples/activities/) - Working with activities

## API Coverage

This client covers all major Strava API v3 endpoints:

- ✅ Activities (CRUD, comments, kudos, laps, zones)
- ✅ Athletes (profile, stats, zones)
- ✅ Clubs (info, members, activities)
- ✅ Segments (leaderboards, efforts, exploration)
- ✅ Routes (info, export GPX/TCX)
- ✅ Gears (bikes and shoes)
- ✅ Streams (time-series data)
- ✅ Uploads (activity file uploads)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This library is distributed under the MIT license. See LICENSE for details.

## Support

For issues and questions, please open an issue on GitHub.

## Disclaimer

This is an unofficial client library and is not affiliated with Strava, Inc.