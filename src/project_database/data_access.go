package project_database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func StartConnection(dbname string, user string, password string) {

	database_connection_arg := user + ":" + password + "@/" + dbname + ""
	db, err := gorm.Open("mysql", database_connection_arg)
	if err != nil {
		log.Fatal(err)
	}
	if !db.HasTable(&User{}) {
		db.CreateTable(&User{})
	}
	if !db.HasTable(&Image{}) {
		db.CreateTable(&User{})
	}
	if !db.HasTable(&Queue{}) {
		db.CreateTable(&User{})
	}
	db.AutoMigrate(&User{}, &Image{}, &Queue{})
}
