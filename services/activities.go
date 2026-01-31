package services

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/kpi-studio/strava-api/models"
)

// ActivitiesService handles activity-related API calls
type ActivitiesService struct {
	client Client
}

// NewActivitiesService creates a new activities service
func NewActivitiesService(client Client) *ActivitiesService {
	return &ActivitiesService{client: client}
}

// List returns a list of activities for the authenticated athlete
func (s *ActivitiesService) List(ctx context.Context, opts *models.ListOptions) ([]*models.Activity, error) {
	path := "/athlete/activities"

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

// Get returns a detailed activity by ID
func (s *ActivitiesService) Get(ctx context.Context, activityID int64, includeAllEfforts bool) (*models.Activity, error) {
	path := fmt.Sprintf("/activities/%d", activityID)

	query := url.Values{}
	if includeAllEfforts {
		query.Set("include_all_efforts", "true")
	}

	var activity models.Activity
	err := s.client.Get(ctx, path, query, &activity)
	return &activity, err
}

// Create creates a new manual activity
func (s *ActivitiesService) Create(ctx context.Context, params models.CreateActivityParams) (*models.Activity, error) {
	path := "/activities"

	data := url.Values{}
	data.Set("name", params.Name)
	data.Set("type", string(params.Type))
	data.Set("start_date_local", params.StartDateLocal.Format("2006-01-02T15:04:05Z"))
	data.Set("elapsed_time", strconv.Itoa(params.ElapsedTime))

	if params.SportType != "" {
		data.Set("sport_type", string(params.SportType))
	}
	if params.Description != "" {
		data.Set("description", params.Description)
	}
	if params.Distance > 0 {
		data.Set("distance", fmt.Sprintf("%.2f", params.Distance))
	}
	if params.Trainer {
		data.Set("trainer", "true")
	}
	if params.Commute {
		data.Set("commute", "true")
	}

	var activity models.Activity
	err := s.client.Post(ctx, path, data, &activity)
	return &activity, err
}

// Update updates an existing activity
func (s *ActivitiesService) Update(ctx context.Context, activityID int64, update *models.UpdatableActivity) (*models.Activity, error) {
	path := fmt.Sprintf("/activities/%d", activityID)

	var activity models.Activity
	err := s.client.Put(ctx, path, update, &activity)
	return &activity, err
}

// Delete deletes an activity
func (s *ActivitiesService) Delete(ctx context.Context, activityID int64) error {
	path := fmt.Sprintf("/activities/%d", activityID)
	return s.client.Delete(ctx, path)
}

// ListComments returns comments for an activity
func (s *ActivitiesService) ListComments(ctx context.Context, activityID int64, pagination *models.Pagination) ([]*models.Comment, error) {
	path := fmt.Sprintf("/activities/%d/comments", activityID)

	query := url.Values{}
	if pagination != nil {
		query = pagination.ToQuery()
	}

	var comments []*models.Comment
	err := s.client.Get(ctx, path, query, &comments)
	return comments, err
}

// ListKudos returns kudos for an activity
func (s *ActivitiesService) ListKudos(ctx context.Context, activityID int64, pagination *models.Pagination) ([]*models.Athlete, error) {
	path := fmt.Sprintf("/activities/%d/kudos", activityID)

	query := url.Values{}
	if pagination != nil {
		query = pagination.ToQuery()
	}

	var kudos []*models.Athlete
	err := s.client.Get(ctx, path, query, &kudos)
	return kudos, err
}

// ListLaps returns laps for an activity
func (s *ActivitiesService) ListLaps(ctx context.Context, activityID int64) ([]*models.Lap, error) {
	path := fmt.Sprintf("/activities/%d/laps", activityID)

	var laps []*models.Lap
	err := s.client.Get(ctx, path, nil, &laps)
	return laps, err
}

// GetZones returns the activity zones (heart rate and/or power)
func (s *ActivitiesService) GetZones(ctx context.Context, activityID int64) (*models.ActivityZones, error) {
	path := fmt.Sprintf("/activities/%d/zones", activityID)

	var zones models.ActivityZones
	err := s.client.Get(ctx, path, nil, &zones)
	return &zones, err
}

// ListRelatedActivities returns activities that were matched as being the same activity
func (s *ActivitiesService) ListRelatedActivities(ctx context.Context, activityID int64, pagination *models.Pagination) ([]*models.Activity, error) {
	path := fmt.Sprintf("/activities/%d/related", activityID)

	query := url.Values{}
	if pagination != nil {
		query = pagination.ToQuery()
	}

	var activities []*models.Activity
	err := s.client.Get(ctx, path, query, &activities)
	return activities, err
}

// GetFeed returns the activities of athletes the authenticated athlete is following
func (s *ActivitiesService) GetFeed(ctx context.Context, opts *models.FeedOptions) ([]*models.Activity, error) {
	path := "/activities/following"

	query := url.Values{}
	if opts != nil {
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

// CreateComment adds a comment to an activity
func (s *ActivitiesService) CreateComment(ctx context.Context, activityID int64, text string) (*models.Comment, error) {
	path := fmt.Sprintf("/activities/%d/comments", activityID)

	data := url.Values{}
	data.Set("text", text)

	var comment models.Comment
	err := s.client.Post(ctx, path, data, &comment)
	return &comment, err
}

// GiveKudos gives kudos to an activity
func (s *ActivitiesService) GiveKudos(ctx context.Context, activityID int64) error {
	path := fmt.Sprintf("/activities/%d/kudos", activityID)
	return s.client.Post(ctx, path, nil, nil)
}
