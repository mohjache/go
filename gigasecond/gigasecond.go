// Package gigasecond has function that adds 1,000,000,000 seconds to a time.
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond takes a time paramater and returns a time 1 000 000 000 seconds later.
func AddGigasecond(t time.Time) time.Time {
	t = t.Add(time.Second * 1000000000)
	return t
}
