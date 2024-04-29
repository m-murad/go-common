package metrics

import "time"

type Service interface {
	Incr(name string, val float64, tags []string)
	Decr(name string, val float64, tags []string)
	Count(name string, val float64, tags []string)
	Timing(name string, dur time.Duration, tags []string)
}
