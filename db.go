package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/servomac/goapi/bundles/todo"
)

type database struct {
	DB *gorm.DB
}

func (db *database) InitDB() {
	var err error
	db.DB, err = gorm.Open("sqlite3", "/tmp/gorm.db")
	if err != nil {
		log.Fatalf("Got an error connecting to database: '%v'", err)
	}
	db.DB.LogMode(true)
}

func (db *database) InitSchema() {
	db.DB.AutoMigrate(&todo.Todo{})
}
