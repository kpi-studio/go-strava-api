# Get Activities Example

This example demonstrates how to fetch a runner's activities from the Strava API.

## Prerequisites

1. A Strava account
2. A Strava API access token

## Getting an Access Token

1. Go to https://www.strava.com/settings/api
2. Create an application if you haven't already
3. Use the authorization flow to get an access token, or use the OAuth example in `../oauth/`

## Running the Example

```bash
# Set your access token
export STRAVA_ACCESS_TOKEN="your_access_token_here"

# Run the example
go run main.go
```

## What It Does

This example:
1. Creates a Strava API client with your access token
2. Fetches the last 10 activities for the authenticated athlete
3. Displays basic information about each activity (name, type, distance, time)
4. Fetches detailed information about the first activity including kudos and comments

## Expected Output

```
Fetching activities...

Found 10 activities:

1. Morning Run
   Type: Run
   Distance: 5.21 km
   Moving Time: 28 minutes
   Date: 2025-12-04 07:30

2. Evening Ride
   Type: Ride
   Distance: 25.50 km
   Moving Time: 65 minutes
   Date: 2025-12-03 18:15

...

Fetching detailed information for first activity...

Detailed Activity: Morning Run
Kudos: 12
Comments: 3
Description: Great morning run!
```
