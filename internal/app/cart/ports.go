package cart

import "github.com/yuraxdrumz/ports-and-adapters-golang/internal/app/cart/structs"

type Port interface {
	Add(item *structs.Item) error
	Remove(itemID string) error
}
