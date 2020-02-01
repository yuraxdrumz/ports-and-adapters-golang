package cartrepository

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
	"github.com/yuraxdrumz/ports-and-adapters-golang/internal/app/cart/structs"
)

// in memory repository adapter, we dont have any dependencies on other adapters so its empty
type SQLiteRepository struct {
	db *sql.DB
}

// new in memory repository factory
func NewSQLiteRepository() *SQLiteRepository {
	db, err := sql.Open("sqlite3", "database/items.db")
	if err != nil {
		logrus.Fatal("Couldn't initialize sqlite db")
	}
	_, err = db.Exec("create table if not exists items (id string, name text, description text)")
	if err != nil {
		logrus.WithField("error", err.Error()).Fatal("Couldn't create initial items table")
	}
	return &SQLiteRepository{
		db: db,
	}
}

// our implementation of the port
func (r *SQLiteRepository) AddItemToDB(item *structs.Item) (string, error) {
	if item.Id == "0" {
		return "", errors.New("cannot add item to db")
	}
	tx, _ := r.db.Begin()
	stmt, _ := tx.Prepare("insert into items (id, name, description) values (?,?,?)")
	_, err := stmt.Exec(item.Id, item.Name, item.Description)
	if err != nil {
		return "", err
	}
	err = tx.Commit()
	if err != nil {
		return "", err
	}
	return "random id", nil
}

// our implementation of the port
func (r *SQLiteRepository) RemoveItemFromDB(itemID string) (bool, error) {
	if itemID == "0" {
		return false, errors.New("item cannot be removed from warehouse, please check again later")
	}
	return true, nil
}
