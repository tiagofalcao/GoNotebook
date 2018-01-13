// Package gigasecond offers gigasecond (1e9 seconds) operations.
package gigasecond

import "time"

// gigasecond represents the duration of 1e9 seconds.
var gigasecond = time.Second * 1e9

// AddGigasecond add 1e9 seconds to time t.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigasecond)
}
