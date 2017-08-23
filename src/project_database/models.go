package project_database

//import "github.com/jinzhu/gorm"

type User struct {
	//gorm.Model
	LOGIN    string `gorm:"size:255"`
	PASSWORD string `gorm:"size:255"`
}

type Image struct {
	//gorm.Model
	PATH string `gorm:"size:255"`
	LABEL string `gorm:"size:600"`
}

type Queue struct {
	//gorm.Model
	UserID []User
	ImageID []Image
}
