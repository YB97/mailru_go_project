package project_database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//"database/sql"
	"log"
	"fmt"
//	"../project_database"

	"encoding/json"
	"os"
)

var db *gorm.DB

type Config struct {
	Database struct{
		Name 	   string  `json:"dbname"`
		User       string  `json:"user"`
		Password   string  `json:"password"`
	} `json:"database"`

	Key string `json:"key"`
}

func LoadConfiguration(file string) Config {
	var config Config
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}

func InitDatabaseConnection(conf Config, user User)  {
	StartConnection(conf.Database.Name, conf.Database.User, conf.Database.Password)
//	fmt.Println()
//	CheckExistAndCreate(conf.Database.Name, conf.Database.User, conf.Database.Password, &user)
}

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

func Get() *gorm.DB {
	return db
}

func CheckExistAndCreate(db *gorm.DB, user *User) {

	//database_connection_arg := username + ":" + password + "@/" + dbname + ""
	//db, err := gorm.Open("mysql", database_connection_arg)
	//if err != nil {
	//	log.Fatal(err)
	//	//return false
	//}

	//user1 := User{LOGIN:"yana", PASSWORD:"1"}
	if db.NewRecord(user) {
		fmt.Println("user")
		db.Create(&user)
	}



//	db.FirstOrCreate(u, &user)
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

