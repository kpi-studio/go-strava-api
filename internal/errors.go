package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Error represents a Strava API error
type Error struct {
	StatusCode int     `json:"status_code"`
	Message    string  `json:"message"`
	Errors     []Fault `json:"errors"`
	Resource   string  `json:"resource"`
	Field      string  `json:"field"`
	Code       string  `json:"code"`
}

// Fault represents a detailed error from the API
type Fault struct {
	Resource string `json:"resource"`
	Field    string `json:"field"`
	Code     string `json:"code"`
}

// Error returns the error message
func (e *Error) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("strava: %s (status: %d)", e.Message, e.StatusCode)
	}
	return fmt.Sprintf("strava: API error (status: %d)", e.StatusCode)
}

// IsRateLimitError checks if the error is a rate limit error
func (e *Error) IsRateLimitError() bool {
	return e.StatusCode == http.StatusTooManyRequests
}

// IsAuthError checks if the error is an authentication error
func (e *Error) IsAuthError() bool {
	return e.StatusCode == http.StatusUnauthorized || e.StatusCode == http.StatusForbidden
}

// IsNotFoundError checks if the error is a not found error
func (e *Error) IsNotFoundError() bool {
	return e.StatusCode == http.StatusNotFound
}

// ParseError parses an error response from the API
func ParseError(resp *http.Response) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &Error{
			StatusCode: resp.StatusCode,
			Message:    "failed to read error response",
		}
	}

	var apiErr Error
	if err := json.Unmarshal(body, &apiErr); err != nil {
		// If we can't parse the error, return a generic one
		return &Error{
			StatusCode: resp.StatusCode,
			Message:    string(body),
		}
	}

	apiErr.StatusCode = resp.StatusCode
	return &apiErr
}

// IsError checks if an error is a Strava API error
func IsError(err error) bool {
	_, ok := err.(*Error)
	return ok
}

// IsRateLimitError checks if an error is a rate limit error
func IsRateLimitError(err error) bool {
	if e, ok := err.(*Error); ok {
		return e.IsRateLimitError()
	}
	return false
}

// IsAuthError checks if an error is an authentication error
func IsAuthError(err error) bool {
	if e, ok := err.(*Error); ok {
		return e.IsAuthError()
	}
	return false
}

// IsNotFoundError checks if an error is a not found error
func IsNotFoundError(err error) bool {
	if e, ok := err.(*Error); ok {
		return e.IsNotFoundError()
	}
	return false
}
