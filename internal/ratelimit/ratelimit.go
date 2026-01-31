package ratelimit

import (
	"context"
	"sync"
	"time"

	"github.com/kpi-studio/strava-api/internal"
)

// RateLimitInfo contains rate limit information from response headers
type RateLimitInfo struct {
	Limit int
	Usage int
	Reset time.Time
}

// RateLimiter manages API rate limiting
type RateLimiter struct {
	mu         sync.Mutex
	limit      int
	usage      int
	reset      time.Time
	minDelay   time.Duration
	maxRetries int
	enabled    bool
}

// RateLimiterConfig contains rate limiter configuration
type RateLimiterConfig struct {
	// Enabled enables rate limiting (default: true)
	Enabled bool

	// MinDelay is the minimum delay between requests (default: 100ms)
	MinDelay time.Duration

	// MaxRetries is the maximum number of retries for rate limited requests (default: 3)
	MaxRetries int
}

// NewRateLimiter creates a new rate limiter
func NewRateLimiter(config *RateLimiterConfig) *RateLimiter {
	rl := &RateLimiter{
		enabled:    true,
		minDelay:   100 * time.Millisecond,
		maxRetries: 3,
	}

	if config != nil {
		if !config.Enabled {
			rl.enabled = false
		}
		if config.MinDelay > 0 {
			rl.minDelay = config.MinDelay
		}
		if config.MaxRetries > 0 {
			rl.maxRetries = config.MaxRetries
		}
	}

	return rl
}

// Wait blocks until it's safe to make another request
func (rl *RateLimiter) Wait(ctx context.Context) error {
	if !rl.enabled {
		return nil
	}

	rl.mu.Lock()
	defer rl.mu.Unlock()

	// Check if we need to wait for rate limit reset
	if rl.limit > 0 && rl.usage >= rl.limit {
		waitTime := time.Until(rl.reset)
		if waitTime > 0 {
			timer := time.NewTimer(waitTime)
			defer timer.Stop()

			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-timer.C:
				// Reset completed, continue
				rl.usage = 0
			}
		}
	}

	// Apply minimum delay between requests
	time.Sleep(rl.minDelay)

	return nil
}

// Update updates rate limit information from response headers
func (rl *RateLimiter) Update(info RateLimitInfo) {
	if !rl.enabled {
		return
	}

	rl.mu.Lock()
	defer rl.mu.Unlock()

	if info.Limit > 0 {
		rl.limit = info.Limit
	}
	if info.Usage > 0 {
		rl.usage = info.Usage
	}
	if !info.Reset.IsZero() {
		rl.reset = info.Reset
	}
}

// RetryWithBackoff retries a function with exponential backoff
func (rl *RateLimiter) RetryWithBackoff(ctx context.Context, fn func() error) error {
	if !rl.enabled {
		return fn()
	}

	var lastErr error
	backoff := 1 * time.Second

	for i := 0; i <= rl.maxRetries; i++ {
		err := fn()
		if err == nil {
			return nil
		}

		lastErr = err

		// Check if it's a rate limit error
		if !internal.IsRateLimitError(err) {
			return err
		}

		// Don't retry on the last attempt
		if i == rl.maxRetries {
			break
		}

		// Wait with exponential backoff
		timer := time.NewTimer(backoff)
		select {
		case <-ctx.Done():
			timer.Stop()
			return ctx.Err()
		case <-timer.C:
		}

		// Exponential backoff with max of 30 seconds
		backoff *= 2
		if backoff > 30*time.Second {
			backoff = 30 * time.Second
		}
	}

	return lastErr
}
