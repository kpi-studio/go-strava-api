package models

import "time"

// Segment represents a Strava segment
type Segment struct {
	ID                  int64                `json:"id"`
	ResourceState       ResourceState        `json:"resource_state"`
	Name                string               `json:"name"`
	ActivityType        ActivityType         `json:"activity_type"`
	Distance            float64              `json:"distance"`
	AverageGrade        float64              `json:"average_grade"`
	MaximumGrade        float64              `json:"maximum_grade"`
	ElevationHigh       float64              `json:"elevation_high"`
	ElevationLow        float64              `json:"elevation_low"`
	StartLatlng         []float64            `json:"start_latlng"`
	EndLatlng           []float64            `json:"end_latlng"`
	ClimbCategory       int                  `json:"climb_category"`
	City                string               `json:"city"`
	State               string               `json:"state"`
	Country             string               `json:"country"`
	Private             bool                 `json:"private"`
	Hazardous           bool                 `json:"hazardous"`
	Starred             bool                 `json:"starred"`
	CreatedAt           time.Time            `json:"created_at"`
	UpdatedAt           time.Time            `json:"updated_at"`
	TotalElevationGain  float64              `json:"total_elevation_gain"`
	Map                 *Map                 `json:"map"`
	EffortCount         int                  `json:"effort_count"`
	AthleteCount        int                  `json:"athlete_count"`
	StarCount           int                  `json:"star_count"`
	AthleteSegmentStats *AthleteSegmentStats `json:"athlete_segment_stats"`
}

// SegmentEffort represents an effort on a segment
type SegmentEffort struct {
	ID               int64         `json:"id"`
	ResourceState    ResourceState `json:"resource_state"`
	Name             string        `json:"name"`
	Activity         *Activity     `json:"activity"`
	Athlete          *Athlete      `json:"athlete"`
	ElapsedTime      float64       `json:"elapsed_time"`
	MovingTime       float64       `json:"moving_time"`
	StartDate        time.Time     `json:"start_date"`
	StartDateLocal   time.Time     `json:"start_date_local"`
	Distance         float64       `json:"distance"`
	StartIndex       int           `json:"start_index"`
	EndIndex         int           `json:"end_index"`
	AverageCadence   float64       `json:"average_cadence"`
	AverageWatts     float64       `json:"average_watts"`
	DeviceWatts      bool          `json:"device_watts"`
	AverageHeartrate float64       `json:"average_heartrate"`
	MaxHeartrate     int           `json:"max_heartrate"`
	Segment          *Segment      `json:"segment"`
	KOMRank          int           `json:"kom_rank"`
	PRRank           int           `json:"pr_rank"`
	Achievements     []Achievement `json:"achievements"`
}

// AthleteSegmentStats represents an athlete's stats on a segment
type AthleteSegmentStats struct {
	PRElapsedTime float64   `json:"pr_elapsed_time"`
	PRDate        time.Time `json:"pr_date"`
	EffortCount   int       `json:"effort_count"`
}

// Leaderboard represents a segment leaderboard
type Leaderboard struct {
	EntryCount  int                 `json:"entry_count"`
	EffortCount int                 `json:"effort_count"`
	KOMType     string              `json:"kom_type"`
	Entries     []*LeaderboardEntry `json:"entries"`
}

// LeaderboardEntry represents an entry in a segment leaderboard
type LeaderboardEntry struct {
	AthleteID      int64     `json:"athlete_id"`
	AthleteName    string    `json:"athlete_name"`
	ElapsedTime    float64   `json:"elapsed_time"`
	MovingTime     float64   `json:"moving_time"`
	StartDate      time.Time `json:"start_date"`
	StartDateLocal time.Time `json:"start_date_local"`
	Rank           int       `json:"rank"`
}

// ExploreResult represents the result of segment exploration
type ExploreResult struct {
	Segments []*ExploreSegment `json:"segments"`
}

// ExploreSegment represents a segment from exploration
type ExploreSegment struct {
	ID                  int64     `json:"id"`
	Name                string    `json:"name"`
	ClimbCategory       int       `json:"climb_category"`
	ClimbCategoryDesc   string    `json:"climb_category_desc"`
	AverageGrade        float64   `json:"avg_grade"`
	StartLatlng         []float64 `json:"start_latlng"`
	EndLatlng           []float64 `json:"end_latlng"`
	ElevationDifference float64   `json:"elev_difference"`
	Distance            float64   `json:"distance"`
	Points              string    `json:"points"`
}

// ListEffortsOptions contains options for listing segment efforts
type ListEffortsOptions struct {
	AthleteID int64
	StartDate time.Time
	EndDate   time.Time
	PerPage   int
}

// LeaderboardOptions contains options for getting a segment leaderboard
type LeaderboardOptions struct {
	Gender         string // M or F
	AgeGroup       string // 0_19, 20_24, 25_34, 35_44, 45_54, 55_64, 65_69, 70_74, 75_plus
	WeightClass    string // 0_124, 125_149, 150_164, 165_179, 180_199, 200_plus (lbs) or 0_54, 55_64, 65_74, 75_84, 85_94, 95_plus (kg)
	Following      bool
	ClubID         int64
	DateRange      string // this_year, this_month, this_week, today
	ContextEntries int
	Page           int
	PerPage        int
}

// ExploreOptions contains options for exploring segments
type ExploreOptions struct {
	Bounds       []float64 // SW lat, SW lng, NE lat, NE lng
	ActivityType string    // riding or running
	MinCat       int
	MaxCat       int
}