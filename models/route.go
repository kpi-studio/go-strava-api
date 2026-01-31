package models

import "time"

// Route represents a Strava route
type Route struct {
	ID                  int64      `json:"id"`
	ResourceState       ResourceState `json:"resource_state"`
	Name                string     `json:"name"`
	Description         string     `json:"description"`
	Athlete             *Athlete   `json:"athlete"`
	Distance            float64    `json:"distance"`
	ElevationGain       float64    `json:"elevation_gain"`
	Map                 *Map       `json:"map"`
	Type                int        `json:"type"`
	SubType             int        `json:"sub_type"`
	Private             bool       `json:"private"`
	Starred             bool       `json:"starred"`
	Timestamp           int        `json:"timestamp"`
	CreatedAt           time.Time  `json:"created_at"`
	UpdatedAt           time.Time  `json:"updated_at"`
	EstimatedMovingTime int        `json:"estimated_moving_time"`
	Waypoints           []Waypoint `json:"waypoints"`
}

// Waypoint represents a waypoint on a route
type Waypoint struct {
	Latlng            []float64 `json:"latlng"`
	TargetLatlng      []float64 `json:"target_latlng"`
	Categories        []string  `json:"categories"`
	Title             string    `json:"title"`
	Description       string    `json:"description"`
	DistanceIntoRoute float64   `json:"distance_into_route"`
}