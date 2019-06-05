package database

import (
    "log"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
    "quarxlab/models"
)

var db *gorm.DB

func init() {

    var err error
    db, err = gorm.Open("sqlite3", "quarxlab.db")
    db.AutoMigrate(&models.Article{})
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
