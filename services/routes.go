package services

import (
	"context"
	"fmt"
	"net/url"

	"github.com/kpi-studio/strava-api/models"
)

// RoutesService handles route-related API calls
type RoutesService struct {
	client Client
}

// NewRoutesService creates a new routes service
func NewRoutesService(client Client) *RoutesService {
	return &RoutesService{client: client}
}

// Get returns a route by ID
func (s *RoutesService) Get(ctx context.Context, routeID int64) (*models.Route, error) {
	path := fmt.Sprintf("/routes/%d", routeID)

	var route models.Route
	err := s.client.Get(ctx, path, nil, &route)
	return &route, err
}

// GetGPX exports a route as GPX
func (s *RoutesService) GetGPX(ctx context.Context, routeID int64) (string, error) {
	path := fmt.Sprintf("/routes/%d/export_gpx", routeID)

	var result struct {
		GPX string `json:"gpx"`
	}
	err := s.client.Get(ctx, path, nil, &result)
	return result.GPX, err
}

// GetTCX exports a route as TCX
func (s *RoutesService) GetTCX(ctx context.Context, routeID int64) (string, error) {
	path := fmt.Sprintf("/routes/%d/export_tcx", routeID)

	var result struct {
		TCX string `json:"tcx"`
	}
	err := s.client.Get(ctx, path, nil, &result)
	return result.TCX, err
}

// ListByAthlete returns routes for an athlete
func (s *RoutesService) ListByAthlete(ctx context.Context, athleteID int64, pagination *models.Pagination) ([]*models.Route, error) {
	path := fmt.Sprintf("/athletes/%d/routes", athleteID)

	query := url.Values{}
	if pagination != nil {
		query = pagination.ToQuery()
	}

	var routes []*models.Route
	err := s.client.Get(ctx, path, query, &routes)
	return routes, err
}