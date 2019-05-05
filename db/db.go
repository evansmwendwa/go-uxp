package db

import (
	"github.com/evansmwendwa/uxp/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // Required for sqlite connection
)

var session *gorm.DB
var err error

func init() {
	db, err := gorm.Open("sqlite3", "data/data.sqlite")
	if err != nil {
		panic("DB connection error")
	}

	db.AutoMigrate(&model.Speaker{})

	session = db
}

// Session - Export a db session for use  by other packages
func Session() *gorm.DB {
	return session
}
