package services

import (
	"context"
	"fmt"

	"github.com/kpi-studio/go-strava-api/models"
)

// UploadsService handles upload-related API calls
type UploadsService struct {
	client Client
}

// NewUploadsService creates a new uploads service
func NewUploadsService(client Client) *UploadsService {
	return &UploadsService{client: client}
}

// GetUploadStatus checks the status of an upload
func (s *UploadsService) GetUploadStatus(ctx context.Context, uploadID int64) (*models.Upload, error) {
	path := fmt.Sprintf("/uploads/%d", uploadID)

	var upload models.Upload
	err := s.client.Get(ctx, path, nil, &upload)
	return &upload, err
}

// TODO: The Upload method needs to be implemented properly
// It requires special handling for multipart form data which isn't
// supported by the simple Client interface
