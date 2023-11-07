package models

import (
	"log"

	"github.com/StefanMoller1/go_library/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func Connect(conf *config.Config, log *log.Logger) (*Database, error) {
	db, err := gorm.Open(postgres.Open(conf.Database.DNS), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &Database{db}, nil
}

func (d *Database) Migrate() error {
	return d.AutoMigrate(&Book{})
}
