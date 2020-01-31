package sleeper

import "time"

// Port - sleep operations
type Port interface {
	Sleep(sec time.Duration)
}
