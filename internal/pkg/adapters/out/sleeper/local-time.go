package sleeper

import (
	"time"
)

// LocalTime - struct with necessary use-cases for adapter to run
type LocalTime struct{}

// NewSleepAdapter - create a new instance of FileUtilsAdapter with passed implementations
func NewSleepAdapter() *LocalTime {
	return &LocalTime{}
}

// Sleep - sleeps for X duration
func (lt *LocalTime) Sleep(sec time.Duration) {
	time.Sleep(sec)
}
