package main

import (
	"context"
	"fmt"
	"log"
	"os"

	strava "github.com/kpi-studio/go-strava-api"
)

func main() {
	// Get access token from environment variable
	accessToken := os.Getenv("STRAVA_ACCESS_TOKEN")
	if accessToken == "" {
		log.Fatal("Please set STRAVA_ACCESS_TOKEN environment variable")
	}

	// Create a new client
	client := strava.NewClient(accessToken)

	// Create context
	ctx := context.Background()

	// Get the authenticated athlete
	athlete, err := client.Athletes.GetCurrent(ctx)
	if err != nil {
		log.Fatalf("Error getting athlete: %v", err)
	}

	fmt.Printf("Hello %s %s!\n", athlete.Firstname, athlete.Lastname)
	fmt.Printf("Athlete ID: %d\n", athlete.ID)
	fmt.Printf("City: %s, %s, %s\n", athlete.City, athlete.State, athlete.Country)

	// List recent activities
	fmt.Println("\n--- Recent Activities ---")
	activities, err := client.Activities.List(ctx, &strava.ListOptions{
		Page:    1,
		PerPage: 10,
	})
	if err != nil {
		log.Fatalf("Error listing activities: %v", err)
	}

	for _, activity := range activities {
		fmt.Printf("%s - %s: %.2f km in %s\n",
			activity.StartDateLocal.Format("Jan 02"),
			activity.Name,
			strava.MetersToKilometers(activity.Distance),
			strava.FormatDuration(activity.MovingTime))
	}

	// Get athlete stats
	fmt.Println("\n--- Athlete Stats ---")
	stats, err := client.Athletes.GetStats(ctx, athlete.ID)
	if err != nil {
		log.Fatalf("Error getting stats: %v", err)
	}

	if stats.YTDRideTotals != nil {
		fmt.Printf("YTD Ride Totals: %d rides, %.2f km, %s\n",
			stats.YTDRideTotals.Count,
			strava.MetersToKilometers(stats.YTDRideTotals.Distance),
			strava.FormatDuration(stats.YTDRideTotals.MovingTime))
	}

	if stats.YTDRunTotals != nil {
		fmt.Printf("YTD Run Totals: %d runs, %.2f km, %s\n",
			stats.YTDRunTotals.Count,
			strava.MetersToKilometers(stats.YTDRunTotals.Distance),
			strava.FormatDuration(stats.YTDRunTotals.MovingTime))
	}
}