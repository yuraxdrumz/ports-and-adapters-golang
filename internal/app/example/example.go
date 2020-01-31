package example

import (
	"time"

	"github.com/yuraxdrumz/golang-starter-kit/internal/pkg/adapters/out/fileutils"
	"github.com/yuraxdrumz/golang-starter-kit/internal/pkg/adapters/out/sleeper"

	log "github.com/sirupsen/logrus"
)

// Something - struct with necessary out adapters
type Something struct {
	checkFileExists fileutils.Port
	sleeper         sleeper.Port
}

// NewExample - create a new instance of Example with passed implementations
func NewExample(checkFileExists fileutils.Port, sleeper sleeper.Port) *Something {
	return &Something{checkFileExists: checkFileExists, sleeper: sleeper}
}

// Run - add method on provided example reference
func (ex *Something) Run() error {
	log.Debug("Running in Example.Run")
	ex.checkFileExists.FileExists("random")
	ex.sleeper.Sleep(time.Duration(10))
	return nil
}
