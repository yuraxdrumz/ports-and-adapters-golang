package cartrepository

import (
	"github.com/yuraxdrumz/ports-and-adapters-golang/internal/app/cart/structs"
)

type Port interface {
	AddItemToDB(item *structs.Item) (string, error)
	RemoveItemFromDB(itemID string) (bool, error)
}
