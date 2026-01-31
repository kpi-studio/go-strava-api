package services

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/kpi-studio/go-strava-api/models"
)

// SegmentsService handles segment-related API calls
type SegmentsService struct {
	client Client
}

// NewSegmentsService creates a new segments service
func NewSegmentsService(client Client) *SegmentsService {
	return &SegmentsService{client: client}
}

// Get returns a segment by ID
func (s *SegmentsService) Get(ctx context.Context, segmentID int64) (*models.Segment, error) {
	path := fmt.Sprintf("/segments/%d", segmentID)

	var segment models.Segment
	err := s.client.Get(ctx, path, nil, &segment)
	return &segment, err
}

// Star stars a segment for the authenticated athlete
func (s *SegmentsService) Star(ctx context.Context, segmentID int64, starred bool) (*models.Segment, error) {
	path := fmt.Sprintf("/segments/%d/starred", segmentID)

	data := url.Values{}
	data.Set("starred", strconv.FormatBool(starred))

	var segment models.Segment
	err := s.client.Put(ctx, path, data, &segment)
	return &segment, err
}

// ListStarred returns the authenticated athlete's starred segments
func (s *SegmentsService) ListStarred(ctx context.Context, pagination *Pagination) ([]*models.Segment, error) {
	path := "/segments/starred"

	query := url.Values{}
	if pagination != nil {
		query = pagination.ToQuery()
	}

	var segments []*models.Segment
	err := s.client.Get(ctx, path, query, &segments)
	return segments, err
}

// GetEffort returns a segment effort by ID
func (s *SegmentsService) GetEffort(ctx context.Context, effortID int64) (*models.SegmentEffort, error) {
	path := fmt.Sprintf("/segment_efforts/%d", effortID)

	var effort models.SegmentEffort
	err := s.client.Get(ctx, path, nil, &effort)
	return &effort, err
}

// ListEfforts returns efforts for a segment
func (s *SegmentsService) ListEfforts(ctx context.Context, segmentID int64, opts *ListEffortsOptions) ([]*models.SegmentEffort, error) {
	path := fmt.Sprintf("/segments/%d/all_efforts", segmentID)

	query := url.Values{}
	if opts != nil {
		if opts.AthleteID > 0 {
			query.Set("athlete_id", strconv.FormatInt(opts.AthleteID, 10))
		}
		if !opts.StartDate.IsZero() {
			query.Set("start_date_local", opts.StartDate.Format("2006-01-02T15:04:05Z"))
		}
		if !opts.EndDate.IsZero() {
			query.Set("end_date_local", opts.EndDate.Format("2006-01-02T15:04:05Z"))
		}
		if opts.PerPage > 0 {
			query.Set("per_page", strconv.Itoa(opts.PerPage))
		}
	}

	var efforts []*models.SegmentEffort
	err := s.client.Get(ctx, path, query, &efforts)
	return efforts, err
}

// ListEffortsOptions contains options for listing segment efforts
type ListEffortsOptions struct {
	AthleteID int64
	StartDate time.Time
	EndDate   time.Time
	PerPage   int
}

// GetLeaderboard returns the leaderboard for a segment
func (s *SegmentsService) GetLeaderboard(ctx context.Context, segmentID int64, opts *LeaderboardOptions) (*models.Leaderboard, error) {
	path := fmt.Sprintf("/segments/%d/leaderboard", segmentID)

	query := url.Values{}
	if opts != nil {
		if opts.Gender != "" {
			query.Set("gender", opts.Gender)
		}
		if opts.AgeGroup != "" {
			query.Set("age_group", opts.AgeGroup)
		}
		if opts.WeightClass != "" {
			query.Set("weight_class", opts.WeightClass)
		}
		if opts.Following {
			query.Set("following", "true")
		}
		if opts.ClubID > 0 {
			query.Set("club_id", strconv.FormatInt(opts.ClubID, 10))
		}
		if opts.DateRange != "" {
			query.Set("date_range", opts.DateRange)
		}
		if opts.ContextEntries > 0 {
			query.Set("context_entries", strconv.Itoa(opts.ContextEntries))
		}
		if opts.Page > 0 {
			query.Set("page", strconv.Itoa(opts.Page))
		}
		if opts.PerPage > 0 {
			query.Set("per_page", strconv.Itoa(opts.PerPage))
		}
	}

	var leaderboard models.Leaderboard
	err := s.client.Get(ctx, path, query, &leaderboard)
	return &leaderboard, err
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
	Bounds      []float64 // SW lat, SW lng, NE lat, NE lng
	ActivityType string   // riding or running
	MinCat      int
	MaxCat      int
}

// Explore finds segments within a given area
func (s *SegmentsService) Explore(ctx context.Context, opts ExploreOptions) (*models.ExploreResult, error) {
	path := "/segments/explore"

	query := url.Values{}
	if len(opts.Bounds) == 4 {
		bounds := []string{
			strconv.FormatFloat(opts.Bounds[0], 'f', -1, 64),
			strconv.FormatFloat(opts.Bounds[1], 'f', -1, 64),
			strconv.FormatFloat(opts.Bounds[2], 'f', -1, 64),
			strconv.FormatFloat(opts.Bounds[3], 'f', -1, 64),
		}
		query.Set("bounds", strings.Join(bounds, ","))
	}
	if opts.ActivityType != "" {
		query.Set("activity_type", opts.ActivityType)
	}
	if opts.MinCat > 0 {
		query.Set("min_cat", strconv.Itoa(opts.MinCat))
	}
	if opts.MaxCat > 0 {
		query.Set("max_cat", strconv.Itoa(opts.MaxCat))
	}

	var result models.ExploreResult
	err := s.client.Get(ctx, path, query, &result)
	return &result, err
}