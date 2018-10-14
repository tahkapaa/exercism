package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond returns the given time + 10^9 seconds
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1e9)
}
