package project_database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)


type Database struct {
	connection  *gorm.DB

}


func (db Database) StartConnection(dbname string, user string, password string) {

	database_connection_arg := user + ":" + password + "@/" + dbname + ""
	conn, err := gorm.Open("mysql", database_connection_arg)
	db.connection = conn

	if err != nil {
		log.Fatal(err)
	}
	if !db.connection.HasTable(&User{}) {
		db.connection.CreateTable(&User{})
	}
	if !db.connection.HasTable(&Image{}) {
		db.connection.CreateTable(&User{})
	}
	if !db.connection.HasTable(&Queue{}) {
		db.connection.CreateTable(&User{})
	}
	db.connection.AutoMigrate(&User{}, &Image{}, &Queue{})
}




func (db Database)  CheckExistance(user User) User {
	err := db.connection.Where(&User{LOGIN: user.LOGIN, PASSWORD: user.PASSWORD}).First(&user)
	if err != nil {
		db.connection.NewRecord(user)
		db.connection.Where(&User{LOGIN: user.LOGIN, PASSWORD: user.PASSWORD}).First(&user)
	}
	return user
}


