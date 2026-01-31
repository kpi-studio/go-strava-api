package services

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/kpi-studio/strava-api/models"
)

// ClubsService handles club-related API calls
type ClubsService struct {
	client Client
}

// NewClubsService creates a new clubs service
func NewClubsService(client Client) *ClubsService {
	return &ClubsService{client: client}
}

// Get returns a club by ID
func (s *ClubsService) Get(ctx context.Context, clubID int64) (*models.Club, error) {
	path := fmt.Sprintf("/clubs/%d", clubID)

	var club models.Club
	err := s.client.Get(ctx, path, nil, &club)
	return &club, err
}

// ListMembers returns members of a club
func (s *ClubsService) ListMembers(ctx context.Context, clubID int64, pagination *Pagination) ([]*models.Athlete, error) {
	path := fmt.Sprintf("/clubs/%d/members", clubID)

	query := url.Values{}
	if pagination != nil {
		query = pagination.ToQuery()
	}

	var members []*models.Athlete
	err := s.client.Get(ctx, path, query, &members)
	return members, err
}

// ListActivities returns activities for a club
func (s *ClubsService) ListActivities(ctx context.Context, clubID int64, opts *models.ListOptions) ([]*models.Activity, error) {
	path := fmt.Sprintf("/clubs/%d/activities", clubID)

	query := url.Values{}
	if opts != nil {
		if opts.Before > 0 {
			query.Set("before", strconv.Itoa(opts.Before))
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

// ListAdmins returns admins of a club
func (s *ClubsService) ListAdmins(ctx context.Context, clubID int64, pagination *Pagination) ([]*models.Athlete, error) {
	path := fmt.Sprintf("/clubs/%d/admins", clubID)

	query := url.Values{}
	if pagination != nil {
		query = pagination.ToQuery()
	}

	var admins []*models.Athlete
	err := s.client.Get(ctx, path, query, &admins)
	return admins, err
}

// ListMyClubs returns clubs the authenticated athlete belongs to
func (s *ClubsService) ListMyClubs(ctx context.Context, pagination *Pagination) ([]*models.Club, error) {
	path := "/athlete/clubs"

	query := url.Values{}
	if pagination != nil {
		query = pagination.ToQuery()
	}

	var clubs []*models.Club
	err := s.client.Get(ctx, path, query, &clubs)
	return clubs, err
}

// Join joins a club
func (s *ClubsService) Join(ctx context.Context, clubID int64) (*models.ClubMembership, error) {
	path := fmt.Sprintf("/clubs/%d/join", clubID)

	var membership models.ClubMembership
	err := s.client.Post(ctx, path, nil, &membership)
	return &membership, err
}

// Leave leaves a club
func (s *ClubsService) Leave(ctx context.Context, clubID int64) (*models.ClubMembership, error) {
	path := fmt.Sprintf("/clubs/%d/leave", clubID)

	var membership models.ClubMembership
	err := s.client.Post(ctx, path, nil, &membership)
	return &membership, err
}

