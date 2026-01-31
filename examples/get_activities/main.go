package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/kpi-studio/strava-api"
	"github.com/kpi-studio/strava-api/models"
)

func main() {
	// Get access token from environment variable
	accessToken := os.Getenv("STRAVA_ACCESS_TOKEN")
	if accessToken == "" {
		log.Fatal("STRAVA_ACCESS_TOKEN environment variable is required.\n\nGet your token from: https://www.strava.com/settings/api\nThen run: export STRAVA_ACCESS_TOKEN=\"your_token_here\"")
	}

	// Create a new Strava client
	client := strava.NewClient(accessToken)

	// Create context
	ctx := context.Background()

	// Get activities with pagination options
	opts := &models.ListOptions{
		Page:    1,
		PerPage: 10,
	}

	fmt.Println("Fetching activities...")
	activities, err := client.Activities.List(ctx, opts)
	if err != nil {
		log.Fatalf("Error fetching activities: %v", err)
	}

	// Display activities
	fmt.Printf("\nFound %d activities:\n\n", len(activities))
	for i, activity := range activities {
		fmt.Printf("%d. %s\n", i+1, activity.Name)
		fmt.Printf("   Type: %s\n", activity.Type)
		fmt.Printf("   Distance: %.2f km\n", activity.Distance/1000)
		fmt.Printf("   Moving Time: %d minutes\n", activity.MovingTime/60)
		fmt.Printf("   Date: %s\n", activity.StartDateLocal.Format("2006-01-02 15:04"))
		fmt.Println()
	}

	// Get detailed information for the first activity
	if len(activities) > 0 {
		fmt.Println("Fetching detailed information for first activity...")
		detailedActivity, err := client.Activities.Get(ctx, activities[0].ID, false)
		if err != nil {
			log.Fatalf("Error fetching activity details: %v", err)
		}

		fmt.Printf("\nDetailed Activity: %s\n", detailedActivity.Name)
		fmt.Printf("Kudos: %d\n", detailedActivity.KudosCount)
		fmt.Printf("Comments: %d\n", detailedActivity.CommentCount)
		if detailedActivity.Description != "" {
			fmt.Printf("Description: %s\n", detailedActivity.Description)
		}
	}
}
