package fileutils

// Port - file operations
type Port interface {
	FileExists(filename string) bool
}
