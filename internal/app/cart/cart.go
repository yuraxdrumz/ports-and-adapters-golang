package cart

import (
	"github.com/yuraxdrumz/ports-and-adapters-golang/internal/app/cart/structs"
	"github.com/yuraxdrumz/ports-and-adapters-golang/internal/pkg/adapters/out/cartRepository"
	"github.com/yuraxdrumz/ports-and-adapters-golang/internal/pkg/adapters/out/warehouse"
	"sync"
)

type Cart struct {
	wh    warehouse.Port
	repo  cartrepository.Port
	mutex sync.Mutex
}

func NewCart(wh warehouse.Port, repo cartrepository.Port) *Cart {
	return &Cart{
		wh:   wh,
		repo: repo,
	}
}

func (c *Cart) Add(item *structs.Item) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	isAvailable, err := c.wh.CheckIfAvailable(item.Id)
	if err != nil {
		return err
	}
	if isAvailable {
		_, err = c.repo.AddItemToDB(item)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Cart) Remove(itemID string) error {
	return nil
}
