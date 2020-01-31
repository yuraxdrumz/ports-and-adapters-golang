package fileutils

import (
	"os"
)

// OsExist - struct with necessary use-cases for adapter to run
type OsExist struct{}

// NewFileUtilsAdapter - create a new instance of FileUtilsAdapter with passed implementations
func NewFileUtilsAdapter() *OsExist {
	return &OsExist{}
}

// FileExists checks if file exists and is not directory
func (dep *OsExist) FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
