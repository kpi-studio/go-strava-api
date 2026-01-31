package services

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/kpi-studio/strava-api/models"
)

// AthletesService handles athlete-related API calls
type AthletesService struct {
	client Client
}

// NewAthletesService creates a new athletes service
func NewAthletesService(client Client) *AthletesService {
	return &AthletesService{client: client}
}

// GetCurrent returns the authenticated athlete
func (s *AthletesService) GetCurrent(ctx context.Context) (*models.Athlete, error) {
	path := "/athlete"

	var athlete models.Athlete
	err := s.client.Get(ctx, path, nil, &athlete)
	return &athlete, err
}

// Get returns an athlete by ID
func (s *AthletesService) Get(ctx context.Context, athleteID int64) (*models.Athlete, error) {
	path := fmt.Sprintf("/athletes/%d", athleteID)

	var athlete models.Athlete
	err := s.client.Get(ctx, path, nil, &athlete)
	return &athlete, err
}

// UpdateWeight updates the authenticated athlete's weight
func (s *AthletesService) UpdateWeight(ctx context.Context, weight float64) (*models.Athlete, error) {
	path := "/athlete"

	data := map[string]interface{}{
		"weight": weight,
	}

	var athlete models.Athlete
	err := s.client.Put(ctx, path, data, &athlete)
	return &athlete, err
}

// GetStats returns statistics for an athlete
func (s *AthletesService) GetStats(ctx context.Context, athleteID int64) (*models.Stats, error) {
	path := fmt.Sprintf("/athletes/%d/stats", athleteID)

	var stats models.Stats
	err := s.client.Get(ctx, path, nil, &stats)
	return &stats, err
}

// ListZones returns the authenticated athlete's heart rate and power zones
func (s *AthletesService) ListZones(ctx context.Context) (*models.AthleteZones, error) {
	path := "/athlete/zones"

	var zones models.AthleteZones
	err := s.client.Get(ctx, path, nil, &zones)
	return &zones, err
}

// ListActivities returns activities for an athlete
func (s *AthletesService) ListActivities(ctx context.Context, athleteID int64, opts *models.ListOptions) ([]*models.Activity, error) {
	path := fmt.Sprintf("/athletes/%d/activities", athleteID)

	query := url.Values{}
	if opts != nil {
		if opts.Before > 0 {
			query.Set("before", strconv.Itoa(opts.Before))
		}
		if opts.After > 0 {
			query.Set("after", strconv.Itoa(opts.After))
		}
		if opts.Page > 0 {
			query.Set("page", strconv.Itoa(opts.Page))
		}
		if opts.PerPage > 0 {
			query.Set("per_page", strconv.Itoa(opts.PerPage))
		}
	}

	var activities []*models.Activity
	err := s.client.Get(ctx, path, query, &activities)
	return activities, err
}

// ListKOMs returns the authenticated athlete's KOMs (King of the Mountains)
func (s *AthletesService) ListKOMs(ctx context.Context, athleteID int64, opts *models.ListKOMsOptions) ([]*models.SegmentEffort, error) {
	path := fmt.Sprintf("/athletes/%d/koms", athleteID)

	query := url.Values{}
	if opts != nil {
		if opts.Page > 0 {
			query.Set("page", strconv.Itoa(opts.Page))
		}
		if opts.PerPage > 0 {
			query.Set("per_page", strconv.Itoa(opts.PerPage))
		}
	}

	var efforts []*models.SegmentEffort
	err := s.client.Get(ctx, path, query, &efforts)
	return efforts, err
}

// ListRoutes returns routes created by the authenticated athlete
func (s *AthletesService) ListRoutes(ctx context.Context, athleteID int64, pagination *models.Pagination) ([]*models.Route, error) {
	path := fmt.Sprintf("/athletes/%d/routes", athleteID)

	query := url.Values{}
	if pagination != nil {
		query = pagination.ToQuery()
	}

	var routes []*models.Route
	err := s.client.Get(ctx, path, query, &routes)
	return routes, err
}
