package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

var db *gorm.DB

func init() {

	var err error
	db, err = gorm.Open("sqlite3", "quarxlab.db")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(20)
	err = db.DB().Ping()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

func Database() *gorm.DB {
	return db
}
