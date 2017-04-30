package todo

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Database struct {
	DB *gorm.DB
}

func (db *Database) InitDB(database string) {
	var err error
	db.DB, err = gorm.Open("sqlite3", database)
	if err != nil {
		log.Fatalf("Error connecting to database: '%v'", err)
	}
	db.DB.LogMode(true)
}

func (db *Database) InitSchema() {
	db.DB.AutoMigrate(&Todo{})
}
