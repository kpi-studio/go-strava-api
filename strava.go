package strava

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/kpi-studio/strava-api/internal"
	"github.com/kpi-studio/strava-api/internal/auth"
	"github.com/kpi-studio/strava-api/internal/ratelimit"
	"github.com/kpi-studio/strava-api/services"
)

const (
	// BaseURL is the default base URL for the Strava API
	BaseURL = "https://www.strava.com/api/v3"

	// DefaultTimeout is the default timeout for HTTP requests
	DefaultTimeout = 30 * time.Second
)

// Client is the main Strava API client
type Client struct {
	httpClient  *http.Client
	baseURL     string
	accessToken string

	// Rate limiter
	rateLimiter *ratelimit.RateLimiter

	// Services
	Activities *services.ActivitiesService
	Athletes   *services.AthletesService
	Clubs      *services.ClubsService
	Segments   *services.SegmentsService
	Routes     *services.RoutesService
	Gears      *services.GearsService
	Streams    *services.StreamsService
	Uploads    *services.UploadsService
}

// ClientOptions contains optional configuration for the client
type ClientOptions struct {
	HTTPClient *http.Client
	BaseURL    string
	RateLimit  *ratelimit.RateLimiterConfig
}

// NewClient creates a new Strava API client with the given access token
func NewClient(accessToken string) *Client {
	return NewClientWithOptions(accessToken, ClientOptions{})
}

// NewClientWithOptions creates a new Strava API client with custom options
func NewClientWithOptions(accessToken string, opts ClientOptions) *Client {
	if opts.HTTPClient == nil {
		opts.HTTPClient = &http.Client{
			Timeout: DefaultTimeout,
		}
	}

	if opts.BaseURL == "" {
		opts.BaseURL = BaseURL
	}

	c := &Client{
		httpClient:  opts.HTTPClient,
		baseURL:     opts.BaseURL,
		accessToken: accessToken,
		rateLimiter: ratelimit.NewRateLimiter(opts.RateLimit),
	}

	// Initialize services with client reference
	c.Activities = services.NewActivitiesService(c)
	c.Athletes = services.NewAthletesService(c)
	c.Clubs = services.NewClubsService(c)
	c.Segments = services.NewSegmentsService(c)
	c.Routes = services.NewRoutesService(c)
	c.Gears = services.NewGearsService(c)
	c.Streams = services.NewStreamsService(c)
	c.Uploads = services.NewUploadsService(c)

	return c
}

// SetAccessToken updates the access token for the client
func (c *Client) SetAccessToken(token string) {
	c.accessToken = token
}

// Response represents an API response
type Response struct {
	*http.Response
	RateLimit ratelimit.RateLimitInfo
}

// NewRequest creates a new API request
func (c *Client) NewRequest(ctx context.Context, method, path string, body interface{}) (*http.Request, error) {
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, err
	}

	var bodyReader io.Reader
	if body != nil {
		switch v := body.(type) {
		case url.Values:
			bodyReader = strings.NewReader(v.Encode())
		case io.Reader:
			bodyReader = v
		default:
			jsonBody, err := json.Marshal(body)
			if err != nil {
				return nil, err
			}
			bodyReader = bytes.NewReader(jsonBody)
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), bodyReader)
	if err != nil {
		return nil, err
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer "+c.accessToken)
	if body != nil {
		switch body.(type) {
		case url.Values:
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		default:
			req.Header.Set("Content-Type", "application/json")
		}
	}

	return req, nil
}

// Do performs an API request
func (c *Client) Do(ctx context.Context, req *http.Request, result interface{}) (*Response, error) {
	// Apply rate limiting
	if err := c.rateLimiter.Wait(ctx); err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response := &Response{
		Response:  resp,
		RateLimit: parseRateLimitHeaders(resp.Header),
	}

	// Update rate limiter with response headers
	c.rateLimiter.Update(response.RateLimit)

	// Check for errors
	if resp.StatusCode >= 400 {
		return response, internal.ParseError(resp)
	}

	// Parse response body if needed
	if result != nil {
		if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
			return response, err
		}
	}

	return response, nil
}

// Get performs a GET request
func (c *Client) Get(ctx context.Context, path string, query url.Values, result interface{}) error {
	if query != nil && len(query) > 0 {
		path = path + "?" + query.Encode()
	}

	req, err := c.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return err
	}

	_, err = c.Do(ctx, req, result)
	return err
}

// Post performs a POST request
func (c *Client) Post(ctx context.Context, path string, body interface{}, result interface{}) error {
	req, err := c.NewRequest(ctx, http.MethodPost, path, body)
	if err != nil {
		return err
	}

	_, err = c.Do(ctx, req, result)
	return err
}

// Put performs a PUT request
func (c *Client) Put(ctx context.Context, path string, body interface{}, result interface{}) error {
	req, err := c.NewRequest(ctx, http.MethodPut, path, body)
	if err != nil {
		return err
	}

	_, err = c.Do(ctx, req, result)
	return err
}

// Delete performs a DELETE request
func (c *Client) Delete(ctx context.Context, path string) error {
	req, err := c.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return err
	}

	_, err = c.Do(ctx, req, nil)
	return err
}

// parseRateLimitHeaders extracts rate limit information from response headers
func parseRateLimitHeaders(headers http.Header) ratelimit.RateLimitInfo {
	info := ratelimit.RateLimitInfo{}

	// Parse rate limit headers
	if limit := headers.Get("X-RateLimit-Limit"); limit != "" {
		fmt.Sscanf(limit, "%d", &info.Limit)
	}

	if usage := headers.Get("X-RateLimit-Usage"); usage != "" {
		fmt.Sscanf(usage, "%d", &info.Usage)
	}

	return info
}

// Re-export auth types for convenience
type (
	OAuth2Config           = auth.OAuth2Config
	TokenResponse          = auth.TokenResponse
	AuthorizationURLParams = auth.AuthorizationURLParams
	TokenManager           = auth.TokenManager
	Scopes                 = auth.Scopes
)

// Re-export auth functions
var (
	NewTokenManager = auth.NewTokenManager
	ParseScopes     = auth.ParseScopes
)
