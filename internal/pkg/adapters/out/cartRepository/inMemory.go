package cartrepository

import (
	"errors"
	"github.com/yuraxdrumz/ports-and-adapters-golang/internal/app/cart/structs"
)

// in memory repository adapter, we dont have any dependencies on other adapters so its empty
type InMemoryRepository struct{}

// new in memory repository factory
func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{}
}

// our implementation of the port
func (r *InMemoryRepository) AddItemToDB(item *structs.Item) (string, error) {
	if item.Id == "0" {
		return "", errors.New("cannot add item to db")
	}
	return "random id", nil
}

// our implementation of the port
func (r *InMemoryRepository) RemoveItemFromDB(itemID string) (bool, error) {
	if itemID == "0" {
		return false, errors.New("item cannot be removed from warehouse, please check again later")
	}
	return true, nil
}
