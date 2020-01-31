package reverser

// InMemory - struct with necessary use-cases for adapter to run
type InMemory struct{}

// NewReverserAdapter - create a new instance of FileUtilsAdapter with passed implementations
func NewReverserAdapter() *InMemory {
	return &InMemory{}
}

// Reverse - reverse bytes
func (in *InMemory) Reverse(byteArray []byte) {
	for i, j := 0, len(byteArray)-1; i < j; i, j = i+1, j-1 {
		byteArray[i], byteArray[j] = byteArray[j], byteArray[i]
	}
}
