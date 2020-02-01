package driver

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"github.com/yuraxdrumz/ports-and-adapters-golang/internal/app/cart"
	"github.com/yuraxdrumz/ports-and-adapters-golang/internal/app/cart/structs"
	"os"
)

// CliAdapter - struct with necessary use-cases for adapter to run
type CliAdapter struct {
	ca cart.Port
}

// NewCliAdapter - create a new instance of NewCliAdapter with passed implementations
func NewCliAdapter(ca cart.Port) *CliAdapter {
	return &CliAdapter{ca: ca}
}

// Run - initializes cli adapter run
func (in *CliAdapter) Run() {
	app := &cli.App{
		Name:  "cart",
		Usage: "handle cart from cli",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "item",
				Usage:    "item to add",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "description",
				Usage:    "description for the item",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "id",
				Usage:    "item id",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {

			item := &structs.Item{
				Name:        c.String("item"),
				Id:          c.String("id"),
				Description: c.String("description"),
			}
			err := in.ca.Add(item)
			if err != nil {
				logrus.WithField("error", err.Error()).Error("couldn't add item")
				return nil
			}
			logrus.WithFields(logrus.Fields{
				"name":        item.Name,
				"description": item.Description,
				"id":          item.Id,
			}).Info("Added new item")
			return nil
		},
	}
	err := app.Run(os.Args)

	if err != nil {
		logrus.Fatal(err)
	}
}
