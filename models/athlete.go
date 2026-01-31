package models

import "time"

// Athlete represents a Strava athlete
type Athlete struct {
	ID            int64         `json:"id"`
	ResourceState ResourceState `json:"resource_state"`
	Firstname     string        `json:"firstname"`
	Lastname      string        `json:"lastname"`
	Username      string        `json:"username"`
	ProfileMedium string        `json:"profile_medium"`
	Profile       string        `json:"profile"`
	City          string        `json:"city"`
	State         string        `json:"state"`
	Country       string        `json:"country"`
	Sex           string        `json:"sex"`
	Premium       bool          `json:"premium"`
	Summit        bool          `json:"summit"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
	FollowerCount int           `json:"follower_count"`
	FriendCount   int           `json:"friend_count"`
	Weight        float64       `json:"weight"` // in kilograms
	Bikes         []Gear        `json:"bikes"`
	Shoes         []Gear        `json:"shoes"`
	Clubs         []Club        `json:"clubs"`
	FTP           int           `json:"ftp"`
}

// Stats represents athlete stats
type Stats struct {
	RecentRideTotals *StatTotals `json:"recent_ride_totals"`
	RecentRunTotals  *StatTotals `json:"recent_run_totals"`
	RecentSwimTotals *StatTotals `json:"recent_swim_totals"`
	YTDRideTotals    *StatTotals `json:"ytd_ride_totals"`
	YTDRunTotals     *StatTotals `json:"ytd_run_totals"`
	YTDSwimTotals    *StatTotals `json:"ytd_swim_totals"`
	AllRideTotals    *StatTotals `json:"all_ride_totals"`
	AllRunTotals     *StatTotals `json:"all_run_totals"`
	AllSwimTotals    *StatTotals `json:"all_swim_totals"`
}

// StatTotals represents statistical totals
type StatTotals struct {
	Count            int     `json:"count"`
	Distance         float64 `json:"distance"`
	MovingTime       int     `json:"moving_time"`
	ElapsedTime      int     `json:"elapsed_time"`
	ElevationGain    float64 `json:"elevation_gain"`
	AchievementCount int     `json:"achievement_count"`
}

// AthleteZones represents an athlete's heart rate and power zones
type AthleteZones struct {
	HeartRate *HeartRateZone `json:"heart_rate"`
	Power     *PowerZone     `json:"power"`
}