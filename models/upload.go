package models

import "io"

// Upload represents an activity upload
type Upload struct {
	ID         int64  `json:"id"`
	ExternalID string `json:"external_id"`
	Error      string `json:"error"`
	Status     string `json:"status"`
	ActivityID int64  `json:"activity_id"`
}

// UploadOptions contains options for uploading an activity file
type UploadOptions struct {
	File         io.Reader
	Name         string
	Description  string
	ActivityType ActivityType
	DataType     string // fit, fit.gz, tcx, tcx.gz, gpx, gpx.gz
	ExternalID   string
	Trainer      bool
	Commute      bool
	Private      bool
}