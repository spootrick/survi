package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spootrick/survi/config"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(config.DBDriver, config.DBUrl)
	if err != nil {
		return nil, err
	}

	return db, nil
}
