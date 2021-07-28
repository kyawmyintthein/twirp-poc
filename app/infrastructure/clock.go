package infrastructure

import "time"

type (
	Clock interface {
		Now() time.Time
		After(d time.Duration) <-chan time.Time
	}
	realClock struct{}
)

func NewClock() Clock {
	return &realClock{}
}
func (realClock) Now() time.Time                         { return time.Now() }
func (realClock) After(d time.Duration) <-chan time.Time { return time.After(d) }
