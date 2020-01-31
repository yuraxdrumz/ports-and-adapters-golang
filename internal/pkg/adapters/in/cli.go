package inadapter

import "github.com/yuraxdrumz/golang-starter-kit/internal/app/example"

// CliAdapter - struct with necessary use-cases for adapter to run
type CliAdapter struct {
	example example.Port
}

// NewCliAdapter - create a new instance of NewCliAdapter with passed implementations
func NewCliAdapter(example example.Port) *CliAdapter {
	return &CliAdapter{example: example}
}

// Run - initializes cli adapter run
func (in *CliAdapter) Run() {
	in.example.Run()
}
