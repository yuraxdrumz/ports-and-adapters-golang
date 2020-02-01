package warehouse

type Port interface {
	CheckIfAvailable(itemID string) (bool, error)
	RemoveItemFromWarehouse(itemID string) (bool, error)
}
