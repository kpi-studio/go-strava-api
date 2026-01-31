package models

import (
	"fmt"
	"net/url"
)

// ActivityType represents the type of activity
type ActivityType string

const (
	ActivityTypeAlpineSki       ActivityType = "AlpineSki"
	ActivityTypeBackcountrySki  ActivityType = "BackcountrySki"
	ActivityTypeCanoeing        ActivityType = "Canoeing"
	ActivityTypeCrossfit        ActivityType = "Crossfit"
	ActivityTypeEBikeRide       ActivityType = "EBikeRide"
	ActivityTypeElliptical      ActivityType = "Elliptical"
	ActivityTypeGolf            ActivityType = "Golf"
	ActivityTypeHandcycle       ActivityType = "Handcycle"
	ActivityTypeHike            ActivityType = "Hike"
	ActivityTypeIceSkate        ActivityType = "IceSkate"
	ActivityTypeInlineSkate     ActivityType = "InlineSkate"
	ActivityTypeKayaking        ActivityType = "Kayaking"
	ActivityTypeKitesurf        ActivityType = "Kitesurf"
	ActivityTypeNordicSki       ActivityType = "NordicSki"
	ActivityTypeRide            ActivityType = "Ride"
	ActivityTypeRockClimbing    ActivityType = "RockClimbing"
	ActivityTypeRollerSki       ActivityType = "RollerSki"
	ActivityTypeRowing          ActivityType = "Rowing"
	ActivityTypeRun             ActivityType = "Run"
	ActivityTypeSail            ActivityType = "Sail"
	ActivityTypeSkateboard      ActivityType = "Skateboard"
	ActivityTypeSnowboard       ActivityType = "Snowboard"
	ActivityTypeSnowshoe        ActivityType = "Snowshoe"
	ActivityTypeStandUpPaddling ActivityType = "StandUpPaddling"
	ActivityTypeSurf            ActivityType = "Surf"
	ActivityTypeSwim            ActivityType = "Swim"
	ActivityTypeVirtualRide     ActivityType = "VirtualRide"
	ActivityTypeVirtualRun      ActivityType = "VirtualRun"
	ActivityTypeWalk            ActivityType = "Walk"
	ActivityTypeWeightTraining  ActivityType = "WeightTraining"
	ActivityTypeWheelchair      ActivityType = "Wheelchair"
	ActivityTypeWindsurf        ActivityType = "Windsurf"
	ActivityTypeWorkout         ActivityType = "Workout"
	ActivityTypeYoga            ActivityType = "Yoga"
)

// SportType represents the sport type of an activity
type SportType string

// ResourceState represents the level of detail of a resource
type ResourceState int

const (
	ResourceStateMeta    ResourceState = 1
	ResourceStateSummary ResourceState = 2
	ResourceStateDetail  ResourceState = 3
)

// Map represents activity map data
type Map struct {
	ID              string        `json:"id"`
	Polyline        string        `json:"polyline"`
	ResourceState   ResourceState `json:"resource_state"`
	SummaryPolyline string        `json:"summary_polyline"`
}

// Photos represents photos attached to an activity
type Photos struct {
	Primary *Photo `json:"primary"`
	Count   int    `json:"count"`
}

// Photo represents a photo
type Photo struct {
	ID       string    `json:"id"`
	UniqueID string    `json:"unique_id"`
	URLs     PhotoURLs `json:"urls"`
	Source   int       `json:"source"`
}

// PhotoURLs contains URLs for different photo sizes
type PhotoURLs struct {
	Size100 string `json:"100"`
	Size600 string `json:"600"`
}

// Achievement represents an achievement earned on a segment effort
type Achievement struct {
	TypeID int    `json:"type_id"`
	Type   string `json:"type"`
	Rank   int    `json:"rank"`
}

// Zone represents a heart rate or power zone
type Zone struct {
	Min  int `json:"min"`
	Max  int `json:"max"`
	Time int `json:"time"`
}

// ZoneRange represents a zone configuration
type ZoneRange struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

// HeartRateZone represents heart rate zones
type HeartRateZone struct {
	CustomZones bool        `json:"custom_zones"`
	Zones       []ZoneRange `json:"zones"`
}

// PowerZone represents power zones
type PowerZone struct {
	CustomZones bool        `json:"custom_zones"`
	Zones       []ZoneRange `json:"zones"`
}

// Pagination represents pagination parameters
type Pagination struct {
	Page    int `json:"page,omitempty"`
	PerPage int `json:"per_page,omitempty"`
	After   int `json:"after,omitempty"`
	Before  int `json:"before,omitempty"`
}

// ToQuery converts pagination to URL query values
func (p *Pagination) ToQuery() url.Values {
	q := url.Values{}
	if p.Page > 0 {
		q.Set("page", fmt.Sprintf("%d", p.Page))
	}
	if p.PerPage > 0 {
		q.Set("per_page", fmt.Sprintf("%d", p.PerPage))
	}
	if p.After > 0 {
		q.Set("after", fmt.Sprintf("%d", p.After))
	}
	if p.Before > 0 {
		q.Set("before", fmt.Sprintf("%d", p.Before))
	}
	return q
}

// ListOptions contains options for listing activities
type ListOptions struct {
	Before  int // Unix timestamp
	After   int // Unix timestamp
	Page    int
	PerPage int
}

// FeedOptions contains options for the activity feed
type FeedOptions struct {
	Page    int
	PerPage int
}

// ListKOMsOptions contains options for listing KOMs
type ListKOMsOptions struct {
	Page    int
	PerPage int
}