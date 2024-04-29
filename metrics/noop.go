package metrics

import "time"

type noOp struct{}

func NewNoOp() *noOp {
	return &noOp{}
}

func (*noOp) Incr(name string, val float64, tags []string)         {}
func (*noOp) Decr(name string, val float64, tags []string)         {}
func (*noOp) Count(name string, val float64, tags []string)        {}
func (*noOp) Timing(name string, dur time.Duration, tags []string) {}
