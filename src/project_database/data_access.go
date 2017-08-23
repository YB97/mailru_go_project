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


func CheckExistAndCreate(dbname string, username string, password string, user *User) {

	database_connection_arg := username + ":" + password + "@/" + dbname + ""
	db, err := gorm.Open("mysql", database_connection_arg)
	if err != nil {
		log.Fatal(err)
		//return false
	}

	//user1 := User{LOGIN:"yana", PASSWORD:"1"}
	if db.NewRecord(user) {
		//db.Create(&user)
	}
	//fmt.Println(a)
	//a:= db.Find(&user, User{LOGIN:"log"})
	//
	//fmt.Println("Checking...")
	//fmt.Println(a.RowsAffected)
	//if a.RowsAffected == 0 {
	//	fmt.Println("yeah")
	//}
	//fmt.Println(a)
	//return user
	//return a
}