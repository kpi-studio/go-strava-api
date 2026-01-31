package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	strava "github.com/kpi-studio/strava-api"
)

func main() {
	// Get access token from environment variable
	accessToken := os.Getenv("STRAVA_ACCESS_TOKEN")
	if accessToken == "" {
		log.Fatal("Please set STRAVA_ACCESS_TOKEN environment variable")
	}

	// Create a new client with custom options
	client := strava.NewClientWithOptions(accessToken, strava.ClientOptions{
		RateLimit: &strava.RateLimiterConfig{
			Enabled:    true,
			MinDelay:   200 * time.Millisecond,
			MaxRetries: 5,
		},
	})

	ctx := context.Background()

	// Example 1: Create a manual activity
	fmt.Println("--- Creating Manual Activity ---")
	newActivity, err := client.Activities.Create(ctx, strava.CreateActivityParams{
		Name:           "Morning Run",
		Type:           strava.ActivityTypeRun,
		StartDateLocal: time.Now().Add(-2 * time.Hour),
		ElapsedTime:    1800, // 30 minutes
		Distance:       5000, // 5km
		Description:    "Easy morning run in the park",
	})
	if err != nil {
		log.Printf("Error creating activity: %v", err)
	} else {
		fmt.Printf("Created activity: %s (ID: %d)\n", newActivity.Name, newActivity.ID)
	}

	// Example 2: Get detailed activity with segments
	fmt.Println("\n--- Get Activity Details ---")
	activities, err := client.Activities.List(ctx, &strava.ListOptions{
		Page:    1,
		PerPage: 1,
	})
	if err != nil {
		log.Fatalf("Error listing activities: %v", err)
	}

	if len(activities) > 0 {
		activityID := activities[0].ID

		// Get full activity details
		activity, err := client.Activities.Get(ctx, activityID, true)
		if err != nil {
			log.Printf("Error getting activity: %v", err)
		} else {
			fmt.Printf("Activity: %s\n", activity.Name)
			fmt.Printf("  Type: %s\n", activity.Type)
			fmt.Printf("  Distance: %.2f km\n", strava.MetersToKilometers(activity.Distance))
			fmt.Printf("  Moving Time: %s\n", strava.FormatDuration(activity.MovingTime))
			fmt.Printf("  Elevation Gain: %.0f m\n", activity.TotalElevationGain)
			fmt.Printf("  Average Speed: %.2f km/h\n", strava.MetersPerSecondToKilometersPerHour(activity.AverageSpeed))

			if activity.AverageHeartrate > 0 {
				fmt.Printf("  Average HR: %.0f bpm\n", activity.AverageHeartrate)
			}
			if activity.AverageWatts > 0 {
				fmt.Printf("  Average Power: %.0f W\n", activity.AverageWatts)
			}

			// Show segment efforts if available
			if len(activity.SegmentEfforts) > 0 {
				fmt.Printf("  Segment Efforts: %d\n", len(activity.SegmentEfforts))
				for i, effort := range activity.SegmentEfforts {
					if i >= 3 {
						break // Show only first 3
					}
					fmt.Printf("    - %s: %s (PR Rank: %d)\n",
						effort.Name,
						strava.FormatDuration(effort.ElapsedTime),
						effort.PRRank)
				}
			}
		}

		// Example 3: Get activity laps
		fmt.Println("\n--- Activity Laps ---")
		laps, err := client.Activities.ListLaps(ctx, activityID)
		if err != nil {
			log.Printf("Error getting laps: %v", err)
		} else {
			for i, lap := range laps {
				fmt.Printf("Lap %d: %.2f km in %s (%.2f km/h)\n",
					i+1,
					strava.MetersToKilometers(lap.Distance),
					strava.FormatDuration(lap.MovingTime),
					strava.MetersPerSecondToKilometersPerHour(lap.AverageSpeed))
			}
		}

		// Example 4: Get activity streams
		fmt.Println("\n--- Activity Streams ---")
		streams, err := client.Streams.GetActivityStreams(ctx, activityID,
			[]strava.StreamType{
				strava.StreamTypeTime,
				strava.StreamTypeDistance,
				strava.StreamTypeAltitude,
				strava.StreamTypeHeartrate,
			},
			"low", // resolution: low, medium, or high
		)
		if err != nil {
			log.Printf("Error getting streams: %v", err)
		} else {
			if streams.Time != nil {
				fmt.Printf("  Time points: %d\n", len(streams.Time.Data))
			}
			if streams.Distance != nil {
				fmt.Printf("  Distance points: %d\n", len(streams.Distance.Data))
			}
			if streams.Altitude != nil {
				fmt.Printf("  Altitude points: %d\n", len(streams.Altitude.Data))
				if len(streams.Altitude.Data) > 0 {
					fmt.Printf("    Min altitude: %.0f m\n", minFloat(streams.Altitude.Data))
					fmt.Printf("    Max altitude: %.0f m\n", maxFloat(streams.Altitude.Data))
				}
			}
			if streams.Heartrate != nil {
				fmt.Printf("  Heart rate points: %d\n", len(streams.Heartrate.Data))
				if len(streams.Heartrate.Data) > 0 {
					fmt.Printf("    Min HR: %d bpm\n", minInt(streams.Heartrate.Data))
					fmt.Printf("    Max HR: %d bpm\n", maxInt(streams.Heartrate.Data))
				}
			}
		}

		// Example 5: Update activity
		fmt.Println("\n--- Update Activity ---")
		updatedActivity, err := client.Activities.Update(ctx, activityID, &strava.UpdatableActivity{
			Description: "Updated description: Great workout!",
		})
		if err != nil {
			log.Printf("Error updating activity: %v", err)
		} else {
			fmt.Printf("Updated activity description: %s\n", updatedActivity.Description)
		}

		// Example 6: Activity social features
		fmt.Println("\n--- Activity Social ---")

		// Give kudos
		err = client.Activities.GiveKudos(ctx, activityID)
		if err != nil {
			log.Printf("Error giving kudos: %v", err)
		} else {
			fmt.Println("Kudos given!")
		}

		// Add a comment
		comment, err := client.Activities.CreateComment(ctx, activityID, "Great effort! Keep it up!")
		if err != nil {
			log.Printf("Error adding comment: %v", err)
		} else {
			fmt.Printf("Comment added: %s\n", comment.Text)
		}

		// List kudos
		kudosAthletes, err := client.Activities.ListKudos(ctx, activityID, &strava.Pagination{
			Page:    1,
			PerPage: 10,
		})
		if err != nil {
			log.Printf("Error listing kudos: %v", err)
		} else {
			fmt.Printf("Kudos from %d athletes\n", len(kudosAthletes))
		}
	}

	// Example 7: Get activity feed
	fmt.Println("\n--- Following Activity Feed ---")
	feed, err := client.Activities.GetFeed(ctx, &strava.FeedOptions{
		Page:    1,
		PerPage: 5,
	})
	if err != nil {
		log.Printf("Error getting feed: %v", err)
	} else {
		for _, activity := range feed {
			fmt.Printf("%s by %s %s: %.2f km\n",
				activity.Name,
				activity.Athlete.Firstname,
				activity.Athlete.Lastname,
				strava.MetersToKilometers(activity.Distance))
		}
	}
}

// Helper functions for min/max
func minFloat(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	min := data[0]
	for _, v := range data[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

func maxFloat(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	max := data[0]
	for _, v := range data[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func minInt(data []int) int {
	if len(data) == 0 {
		return 0
	}
	min := data[0]
	for _, v := range data[1:] {
		if v < min {
			min = v
		}
	}
	return min
}

func maxInt(data []int) int {
	if len(data) == 0 {
		return 0
	}
	max := data[0]
	for _, v := range data[1:] {
		if v > max {
			max = v
		}
	}
	return max
}