package todo

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Database struct {
	DB *gorm.DB
}

func (db *Database) InitDB() {
	var err error
	db.DB, err = gorm.Open("sqlite3", "/tmp/gorm.db")
	if err != nil {
		log.Fatalf("Got an error connecting to database: '%v'", err)
	}
	db.DB.LogMode(true)
}

func (db *Database) InitSchema() {
	db.DB.AutoMigrate(&Todo{})
}