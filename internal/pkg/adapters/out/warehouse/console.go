package warehouse

import (
	"errors"
)

// our console warehouse struct with all the adapters it uses, in this case with the console adapter we dont need any dependencies
type ConsoleWarehouse struct{}

// a new console warehouse factory
func NewConsoleWarehouse() *ConsoleWarehouse {
	return &ConsoleWarehouse{}
}

// our implementation of the port for console warehouse
func (w *ConsoleWarehouse) CheckIfAvailable(itemID string) (bool, error) {
	if itemID == "0" {
		return false, errors.New("item is not available")
	}
	return true, nil
}

// our implementation of the port for console warehouse
func (w *ConsoleWarehouse) RemoveItemFromWarehouse(itemID string) (bool, error) {
	if itemID == "0" {
		return false, errors.New("item cannot be removed from warehouse, please check again later")
	}
	return true, nil
}
