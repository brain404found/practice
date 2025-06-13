package limiter

import (
	"context"
	"time"
)

type RateLimiter struct {
	ch chan struct{}
}

// NewRateLimiter creates a new rate limiter with the given limit and interval
// Limit is the maximum number of requests allowed in the given interval
// Interval is the time window for the rate limit. Ex: limit - 10, interval - 1 second, 
// no more than 10 RPS
func NewRateLimiter(ctx context.Context, limit int, interval time.Duration) *RateLimiter {
	limiter := &RateLimiter{
		ch: make(chan struct{}, limit),
	}

	for range limit {
		limiter.ch <- struct{}{}
	}

	tokenDelay := interval.Nanoseconds() / int64(limit)

	go func() {
		ticker := time.NewTicker(time.Duration(tokenDelay))
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				limiter.ch <- struct{}{}
			}
		}
	}()

	return limiter
}

// Allow checks if the request is allowed to proceed
func (l *RateLimiter) Allow() bool {
	select {
	case <-l.ch:
		return true
	default:
		return false
	}
}
