package main

import (
	"github.com/yuraxdrumz/ports-and-adapters-golang/internal/app/cart"
	driver "github.com/yuraxdrumz/ports-and-adapters-golang/internal/pkg/adapters/in"
	"github.com/yuraxdrumz/ports-and-adapters-golang/internal/pkg/adapters/out/cartRepository"
	"github.com/yuraxdrumz/ports-and-adapters-golang/internal/pkg/adapters/out/warehouse"
)

// with http adapter
func main() {
	// declare all ports
	var ca cart.Port
	var wh warehouse.Port
	var re cartrepository.Port
	var in driver.Port

	wh = warehouse.NewConsoleWarehouse()
	//re = cartrepository.NewInMemoryRepository()
	re = cartrepository.NewSQLiteRepository()
	ca = cart.NewCart(wh, re)
	in = driver.NewCliAdapter(ca)

	in.Run()
}
