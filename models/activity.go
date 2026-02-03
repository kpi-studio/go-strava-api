package models

import "time"

// Activity represents a Strava activity
type Activity struct {
	ID                   int64         `json:"id"`
	ResourceState        ResourceState `json:"resource_state"`
	ExternalID           string        `json:"external_id"`
	UploadID             int64         `json:"upload_id"`
	Athlete              *Athlete      `json:"athlete"`
	Name                 string        `json:"name"`
	Distance             float64       `json:"distance"`
	MovingTime           float64       `json:"moving_time"`
	ElapsedTime          float64       `json:"elapsed_time"`
	TotalElevationGain   float64       `json:"total_elevation_gain"`
	Type                 ActivityType  `json:"type"`
	SportType            SportType     `json:"sport_type"`
	StartDate            time.Time     `json:"start_date"`
	StartDateLocal       time.Time     `json:"start_date_local"`
	Timezone             string        `json:"timezone"`
	UTCOffset            float64       `json:"utc_offset"`
	StartLatlng          []float64     `json:"start_latlng"`
	EndLatlng            []float64     `json:"end_latlng"`
	AchievementCount     int           `json:"achievement_count"`
	KudosCount           int           `json:"kudos_count"`
	CommentCount         int           `json:"comment_count"`
	AthleteCount         int           `json:"athlete_count"`
	PhotoCount           int           `json:"photo_count"`
	Map                  *Map          `json:"map"`
	Trainer              bool          `json:"trainer"`
	Commute              bool          `json:"commute"`
	Manual               bool          `json:"manual"`
	Private              bool          `json:"private"`
	Flagged              bool          `json:"flagged"`
	GearID               string        `json:"gear_id"`
	FromAcceptedTag      bool          `json:"from_accepted_tag"`
	AverageSpeed         float64       `json:"average_speed"`
	MaxSpeed             float64       `json:"max_speed"`
	AverageCadence       float64       `json:"average_cadence"`
	AverageTemp          int           `json:"average_temp"`
	AverageWatts         float64       `json:"average_watts"`
	WeightedAverageWatts int           `json:"weighted_average_watts"`
	Kilojoules           float64       `json:"kilojoules"`
	DeviceWatts          bool          `json:"device_watts"`
	HasHeartrate         bool          `json:"has_heartrate"`
	AverageHeartrate     float64       `json:"average_heartrate"`
	MaxHeartrate         int           `json:"max_heartrate"`
	MaxWatts             int           `json:"max_watts"`
	PRCount              int           `json:"pr_count"`
	TotalPhotoCount      int           `json:"total_photo_count"`
	HasKudoed            bool          `json:"has_kudoed"`
	SufferScore          int           `json:"suffer_score"`
	Description          string        `json:"description"`
	Calories             float64       `json:"calories"`
	SegmentEfforts       []SegmentEffort `json:"segment_efforts"`
	SplitsMetric         []Split       `json:"splits_metric"`
	SplitsStandard       []Split       `json:"splits_standard"`
	Laps                 []Lap         `json:"laps"`
	Gear                 *Gear         `json:"gear"`
	Photos               *Photos       `json:"photos"`
	DeviceName           string        `json:"device_name"`
	EmbedToken           string        `json:"embed_token"`
}

// UpdatableActivity represents fields that can be updated for an activity
type UpdatableActivity struct {
	Name         string       `json:"name,omitempty"`
	Type         ActivityType `json:"type,omitempty"`
	SportType    SportType    `json:"sport_type,omitempty"`
	GearID       string       `json:"gear_id,omitempty"`
	Description  string       `json:"description,omitempty"`
	Trainer      *bool        `json:"trainer,omitempty"`
	Commute      *bool        `json:"commute,omitempty"`
	HideFromHome *bool        `json:"hide_from_home,omitempty"`
}

// CreateActivityParams contains parameters for creating an activity
type CreateActivityParams struct {
	Name           string
	Type           ActivityType
	SportType      SportType
	StartDateLocal time.Time
	ElapsedTime    int     // in seconds
	Description    string
	Distance       float64 // in meters
	Trainer        bool
	Commute        bool
}

// Split represents a split (metric or imperial)
type Split struct {
	Distance            float64 `json:"distance"`
	ElapsedTime         float64 `json:"elapsed_time"`
	MovingTime          float64 `json:"moving_time"`
	ElevationDifference float64 `json:"elevation_difference"`
	SplitNumber         int     `json:"split"`
	AverageSpeed        float64 `json:"average_speed"`
	PaceZone            int     `json:"pace_zone"`
}

// Lap represents a lap in an activity
type Lap struct {
	ID                 int64         `json:"id"`
	ResourceState      ResourceState `json:"resource_state"`
	Name               string        `json:"name"`
	Activity           *Activity     `json:"activity"`
	Athlete            *Athlete      `json:"athlete"`
	ElapsedTime        float64       `json:"elapsed_time"`
	MovingTime         float64       `json:"moving_time"`
	StartDate          time.Time     `json:"start_date"`
	StartDateLocal     time.Time     `json:"start_date_local"`
	Distance           float64       `json:"distance"`
	StartIndex         int           `json:"start_index"`
	EndIndex           int           `json:"end_index"`
	TotalElevationGain float64       `json:"total_elevation_gain"`
	AverageSpeed       float64       `json:"average_speed"`
	MaxSpeed           float64       `json:"max_speed"`
	AverageCadence     float64       `json:"average_cadence"`
	DeviceWatts        bool          `json:"device_watts"`
	AverageWatts       float64       `json:"average_watts"`
	AverageHeartrate   float64       `json:"average_heartrate"`
	MaxHeartrate       int           `json:"max_heartrate"`
	LapIndex           int           `json:"lap_index"`
	Split              int           `json:"split"`
	PaceZone           int           `json:"pace_zone"`
}

// Comment represents a comment on an activity
type Comment struct {
	ID         int64     `json:"id"`
	ActivityID int64     `json:"activity_id"`
	Text       string    `json:"text"`
	Athlete    *Athlete  `json:"athlete"`
	CreatedAt  time.Time `json:"created_at"`
}

// Kudos represents kudos given to an activity
type Kudos struct {
	Count    int        `json:"count"`
	Athletes []*Athlete `json:"athletes"`
}

// ActivityZones represents activity zones
type ActivityZones struct {
	Score               int          `json:"score"`
	DistributionBuckets []ZoneBucket `json:"distribution_buckets"`
	Type                string       `json:"type"`
	ResourceState       ResourceState `json:"resource_state"`
	SensorBased         bool         `json:"sensor_based"`
	Points              int          `json:"points"`
	CustomZones         bool         `json:"custom_zones"`
	Max                 int          `json:"max"`
}

// ZoneBucket represents a zone bucket in the distribution
type ZoneBucket struct {
	Min  int `json:"min"`
	Max  int `json:"max"`
	Time int `json:"time"`
}