package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/kpi-studio/go-strava-api/internal"
	"github.com/kpi-studio/go-strava-api/models"
)

// OAuth2Config contains OAuth 2.0 configuration
type OAuth2Config struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
	Scopes       []string
}

// AuthorizationURLParams contains parameters for building the authorization URL
type AuthorizationURLParams struct {
	ApprovalPrompt string // "force" or "auto"
	State          string // Optional state parameter
}

// TokenResponse represents the OAuth token response
type TokenResponse struct {
	TokenType    string          `json:"token_type"`
	ExpiresAt    int64           `json:"expires_at"`
	ExpiresIn    int             `json:"expires_in"`
	RefreshToken string          `json:"refresh_token"`
	AccessToken  string          `json:"access_token"`
	Athlete      *models.Athlete `json:"athlete"`
}

// IsExpired checks if the token is expired
func (t *TokenResponse) IsExpired() bool {
	return time.Now().Unix() >= t.ExpiresAt
}

// ExpirationTime returns the expiration time of the token
func (t *TokenResponse) ExpirationTime() time.Time {
	return time.Unix(t.ExpiresAt, 0)
}

// TimeUntilExpiration returns the duration until the token expires
func (t *TokenResponse) TimeUntilExpiration() time.Duration {
	return time.Until(t.ExpirationTime())
}

// GetAuthorizationURL returns the OAuth authorization URL
func (c *OAuth2Config) GetAuthorizationURL(params AuthorizationURLParams) string {
	u, _ := url.Parse("https://www.strava.com/oauth/authorize")

	q := u.Query()
	q.Set("client_id", c.ClientID)
	q.Set("response_type", "code")
	q.Set("redirect_uri", c.RedirectURI)

	if len(c.Scopes) > 0 {
		q.Set("scope", strings.Join(c.Scopes, ","))
	}

	if params.ApprovalPrompt != "" {
		q.Set("approval_prompt", params.ApprovalPrompt)
	}

	if params.State != "" {
		q.Set("state", params.State)
	}

	u.RawQuery = q.Encode()
	return u.String()
}

// ExchangeCode exchanges an authorization code for an access token
func (c *OAuth2Config) ExchangeCode(ctx context.Context, code string) (*TokenResponse, error) {
	data := url.Values{
		"client_id":     {c.ClientID},
		"client_secret": {c.ClientSecret},
		"code":          {code},
		"grant_type":    {"authorization_code"},
	}

	return c.doTokenRequest(ctx, data)
}

// RefreshToken refreshes an access token using a refresh token
func (c *OAuth2Config) RefreshToken(ctx context.Context, refreshToken string) (*TokenResponse, error) {
	data := url.Values{
		"client_id":     {c.ClientID},
		"client_secret": {c.ClientSecret},
		"refresh_token": {refreshToken},
		"grant_type":    {"refresh_token"},
	}

	return c.doTokenRequest(ctx, data)
}

// doTokenRequest performs the token request
func (c *OAuth2Config) doTokenRequest(ctx context.Context, data url.Values) (*TokenResponse, error) {
	req, err := http.NewRequestWithContext(ctx, "POST", "https://www.strava.com/oauth/token", strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, internal.ParseError(resp)
	}

	var tokenResp TokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
		return nil, fmt.Errorf("failed to decode token response: %w", err)
	}

	return &tokenResp, nil
}

// TokenManager manages token lifecycle with automatic refresh
type TokenManager struct {
	config       *OAuth2Config
	token        *TokenResponse
	refreshToken string
	onTokenUpdate func(*TokenResponse)
}

// NewTokenManager creates a new token manager
func NewTokenManager(config *OAuth2Config, token *TokenResponse) *TokenManager {
	return &TokenManager{
		config:       config,
		token:        token,
		refreshToken: token.RefreshToken,
	}
}

// SetTokenUpdateCallback sets a callback function that will be called when the token is updated
func (tm *TokenManager) SetTokenUpdateCallback(callback func(*TokenResponse)) {
	tm.onTokenUpdate = callback
}

// GetAccessToken returns the current access token, refreshing if necessary
func (tm *TokenManager) GetAccessToken(ctx context.Context) (string, error) {
	if tm.token == nil {
		return "", fmt.Errorf("no token available")
	}

	// Check if token needs refresh (with 5 minute buffer)
	if time.Now().Unix() >= (tm.token.ExpiresAt - 300) {
		if tm.refreshToken == "" {
			return "", fmt.Errorf("token expired and no refresh token available")
		}

		// Refresh the token
		newToken, err := tm.config.RefreshToken(ctx, tm.refreshToken)
		if err != nil {
			return "", fmt.Errorf("failed to refresh token: %w", err)
		}

		tm.token = newToken
		tm.refreshToken = newToken.RefreshToken

		// Call update callback if set
		if tm.onTokenUpdate != nil {
			tm.onTokenUpdate(newToken)
		}
	}

	return tm.token.AccessToken, nil
}

// GetToken returns the current token
func (tm *TokenManager) GetToken() *TokenResponse {
	return tm.token
}

// UpdateToken manually updates the token
func (tm *TokenManager) UpdateToken(token *TokenResponse) {
	tm.token = token
	if token.RefreshToken != "" {
		tm.refreshToken = token.RefreshToken
	}
}

// Scopes represents Strava API scopes
type Scopes struct {
	Read         bool
	ReadAll      bool
	ProfileRead  bool
	ProfileWrite bool
	ActivityRead bool
	ActivityReadAll bool
	ActivityWrite bool
}

// ToStringSlice converts scopes to a string slice
func (s Scopes) ToStringSlice() []string {
	var scopes []string

	if s.Read {
		scopes = append(scopes, "read")
	}
	if s.ReadAll {
		scopes = append(scopes, "read_all")
	}
	if s.ProfileRead {
		scopes = append(scopes, "profile:read_all")
	}
	if s.ProfileWrite {
		scopes = append(scopes, "profile:write")
	}
	if s.ActivityRead {
		scopes = append(scopes, "activity:read")
	}
	if s.ActivityReadAll {
		scopes = append(scopes, "activity:read_all")
	}
	if s.ActivityWrite {
		scopes = append(scopes, "activity:write")
	}

	return scopes
}

// ParseScopes parses a comma-separated string of scopes
func ParseScopes(scopeString string) Scopes {
	s := Scopes{}
	scopes := strings.Split(scopeString, ",")

	for _, scope := range scopes {
		switch strings.TrimSpace(scope) {
		case "read":
			s.Read = true
		case "read_all":
			s.ReadAll = true
		case "profile:read_all":
			s.ProfileRead = true
		case "profile:write":
			s.ProfileWrite = true
		case "activity:read":
			s.ActivityRead = true
		case "activity:read_all":
			s.ActivityReadAll = true
		case "activity:write":
			s.ActivityWrite = true
		}
	}

	return s
}