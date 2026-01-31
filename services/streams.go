package services

import (
	"github.com/kpi-studio/go-strava-api/models"
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

// StreamsService handles stream-related API calls
type StreamsService struct {
	client Client
}

// NewStreamsService creates a new streams service
func NewStreamsService(client Client) *StreamsService {
	return &StreamsService{client: client}
}

// GetActivityStreams returns streams for an activity
func (s *StreamsService) GetActivityStreams(ctx context.Context, activityID int64, types []models.StreamType, resolution string) (*models.StreamSet, error) {
	path := fmt.Sprintf("/activities/%d/streams", activityID)

	// Convert stream types to strings
	typeStrings := make([]string, len(types))
	for i, t := range types {
		typeStrings[i] = string(t)
	}

	query := url.Values{}
	query.Set("keys", strings.Join(typeStrings, ","))
	if resolution != "" {
		query.Set("key_by_type", "true")
		query.Set("resolution", resolution)
	} else {
		query.Set("key_by_type", "true")
	}

	// Get raw response first
	var rawStreams []json.RawMessage
	err := s.client.Get(ctx, path, query, &rawStreams)
	if err != nil {
		return nil, err
	}

	// Parse into StreamSet
	streamSet := &models.StreamSet{}
	for _, raw := range rawStreams {
		var base models.BaseStream
		if err := json.Unmarshal(raw, &base); err != nil {
			continue
		}

		switch models.StreamType(base.Type) {
		case models.StreamTypeTime:
			var stream models.TimeStream
			json.Unmarshal(raw, &stream)
			streamSet.Time = &stream
		case models.StreamTypeDistance:
			var stream models.DistanceStream
			json.Unmarshal(raw, &stream)
			streamSet.Distance = &stream
		case models.StreamTypeLatLng:
			var stream models.LatLngStream
			json.Unmarshal(raw, &stream)
			streamSet.LatLng = &stream
		case models.StreamTypeAltitude:
			var stream models.AltitudeStream
			json.Unmarshal(raw, &stream)
			streamSet.Altitude = &stream
		case models.StreamTypeVelocity:
			var stream models.VelocityStream
			json.Unmarshal(raw, &stream)
			streamSet.VelocitySmooth = &stream
		case models.StreamTypeHeartrate:
			var stream models.HeartrateStream
			json.Unmarshal(raw, &stream)
			streamSet.Heartrate = &stream
		case models.StreamTypeCadence:
			var stream models.CadenceStream
			json.Unmarshal(raw, &stream)
			streamSet.Cadence = &stream
		case models.StreamTypePower:
			var stream models.PowerStream
			json.Unmarshal(raw, &stream)
			streamSet.Watts = &stream
		case models.StreamTypeTemperature:
			var stream models.TemperatureStream
			json.Unmarshal(raw, &stream)
			streamSet.Temperature = &stream
		case models.StreamTypeMoving:
			var stream models.MovingStream
			json.Unmarshal(raw, &stream)
			streamSet.Moving = &stream
		case models.StreamTypeGrade:
			var stream models.GradeStream
			json.Unmarshal(raw, &stream)
			streamSet.GradeSmooth = &stream
		}
	}

	return streamSet, nil
}

// GetSegmentStreams returns streams for a segment
func (s *StreamsService) GetSegmentStreams(ctx context.Context, segmentID int64, types []models.StreamType, resolution string) (*models.StreamSet, error) {
	path := fmt.Sprintf("/segments/%d/streams", segmentID)

	// Convert stream types to strings
	typeStrings := make([]string, len(types))
	for i, t := range types {
		typeStrings[i] = string(t)
	}

	query := url.Values{}
	query.Set("keys", strings.Join(typeStrings, ","))
	if resolution != "" {
		query.Set("key_by_type", "true")
		query.Set("resolution", resolution)
	} else {
		query.Set("key_by_type", "true")
	}

	// Get raw response first
	var rawStreams []json.RawMessage
	err := s.client.Get(ctx, path, query, &rawStreams)
	if err != nil {
		return nil, err
	}

	// Parse into StreamSet
	streamSet := &models.StreamSet{}
	for _, raw := range rawStreams {
		var base models.BaseStream
		if err := json.Unmarshal(raw, &base); err != nil {
			continue
		}

		switch models.StreamType(base.Type) {
		case models.StreamTypeDistance:
			var stream models.DistanceStream
			json.Unmarshal(raw, &stream)
			streamSet.Distance = &stream
		case models.StreamTypeLatLng:
			var stream models.LatLngStream
			json.Unmarshal(raw, &stream)
			streamSet.LatLng = &stream
		case models.StreamTypeAltitude:
			var stream models.AltitudeStream
			json.Unmarshal(raw, &stream)
			streamSet.Altitude = &stream
		}
	}

	return streamSet, nil
}

// GetSegmentEffortStreams returns streams for a segment effort
func (s *StreamsService) GetSegmentEffortStreams(ctx context.Context, effortID int64, types []models.StreamType, resolution string) (*models.StreamSet, error) {
	path := fmt.Sprintf("/segment_efforts/%d/streams", effortID)

	// Convert stream types to strings
	typeStrings := make([]string, len(types))
	for i, t := range types {
		typeStrings[i] = string(t)
	}

	query := url.Values{}
	query.Set("keys", strings.Join(typeStrings, ","))
	if resolution != "" {
		query.Set("key_by_type", "true")
		query.Set("resolution", resolution)
	} else {
		query.Set("key_by_type", "true")
	}

	// Get raw response first
	var rawStreams []json.RawMessage
	err := s.client.Get(ctx, path, query, &rawStreams)
	if err != nil {
		return nil, err
	}

	// Parse into models.StreamSet similar to GetActivityStreams
	streamSet := &models.StreamSet{}
	for _, raw := range rawStreams {
		var base models.BaseStream
		if err := json.Unmarshal(raw, &base); err != nil {
			continue
		}

		// Parse based on type (similar to GetActivityStreams)
		switch models.StreamType(base.Type) {
		case models.StreamTypeTime:
			var stream models.TimeStream
			json.Unmarshal(raw, &stream)
			streamSet.Time = &stream
		case models.StreamTypeDistance:
			var stream models.DistanceStream
			json.Unmarshal(raw, &stream)
			streamSet.Distance = &stream
		case models.StreamTypeLatLng:
			var stream models.LatLngStream
			json.Unmarshal(raw, &stream)
			streamSet.LatLng = &stream
		case models.StreamTypeAltitude:
			var stream models.AltitudeStream
			json.Unmarshal(raw, &stream)
			streamSet.Altitude = &stream
		case models.StreamTypeVelocity:
			var stream models.VelocityStream
			json.Unmarshal(raw, &stream)
			streamSet.VelocitySmooth = &stream
		case models.StreamTypeHeartrate:
			var stream models.HeartrateStream
			json.Unmarshal(raw, &stream)
			streamSet.Heartrate = &stream
		case models.StreamTypeCadence:
			var stream models.CadenceStream
			json.Unmarshal(raw, &stream)
			streamSet.Cadence = &stream
		case models.StreamTypePower:
			var stream models.PowerStream
			json.Unmarshal(raw, &stream)
			streamSet.Watts = &stream
		case models.StreamTypeMoving:
			var stream models.MovingStream
			json.Unmarshal(raw, &stream)
			streamSet.Moving = &stream
		case models.StreamTypeGrade:
			var stream models.GradeStream
			json.Unmarshal(raw, &stream)
			streamSet.GradeSmooth = &stream
		}
	}

	return streamSet, nil
}

// GetRouteStreams returns streams for a route
func (s *StreamsService) GetRouteStreams(ctx context.Context, routeID int64, types []models.StreamType) (*models.StreamSet, error) {
	path := fmt.Sprintf("/routes/%d/streams", routeID)

	// Convert stream types to strings
	typeStrings := make([]string, len(types))
	for i, t := range types {
		typeStrings[i] = string(t)
	}

	query := url.Values{}
	query.Set("keys", strings.Join(typeStrings, ","))
	query.Set("key_by_type", "true")

	// Get raw response first
	var rawStreams []json.RawMessage
	err := s.client.Get(ctx, path, query, &rawStreams)
	if err != nil {
		return nil, err
	}

	// Parse into StreamSet
	streamSet := &models.StreamSet{}
	for _, raw := range rawStreams {
		var base models.BaseStream
		if err := json.Unmarshal(raw, &base); err != nil {
			continue
		}

		switch models.StreamType(base.Type) {
		case models.StreamTypeDistance:
			var stream models.DistanceStream
			json.Unmarshal(raw, &stream)
			streamSet.Distance = &stream
		case models.StreamTypeLatLng:
			var stream models.LatLngStream
			json.Unmarshal(raw, &stream)
			streamSet.LatLng = &stream
		case models.StreamTypeAltitude:
			var stream models.AltitudeStream
			json.Unmarshal(raw, &stream)
			streamSet.Altitude = &stream
		}
	}

	return streamSet, nil
}