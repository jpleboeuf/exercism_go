/*
Package gigasecond implements utility routines
 returning miscellaneous moments. 
*/
package gigasecond

// Import the time package from the standard library.
import "time"

// Gigasecond represents the Duration of a gigasecond,
//  defined as 10^9 (1,000,000,000) Seconds.
const Gigasecond time.Duration = 1e9 * time.Second;

// AddGigasecond returns,
//  given a moment t,
//  the moment that would be after a gigasecond has passed since t.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(Gigasecond)
}
