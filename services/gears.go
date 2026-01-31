package services

import (
	"context"
	"fmt"

	"github.com/kpi-studio/strava-api/models"
)

// GearsService handles gear-related API calls
type GearsService struct {
	client Client
}

// NewGearsService creates a new gears service
func NewGearsService(client Client) *GearsService {
	return &GearsService{client: client}
}

// Get returns gear by ID
func (s *GearsService) Get(ctx context.Context, gearID string) (*models.Gear, error) {
	path := fmt.Sprintf("/gear/%s", gearID)

	var gear models.Gear
	err := s.client.Get(ctx, path, nil, &gear)
	return &gear, err
}