package project_database

type User struct {
	LOGIN    string `gorm:"size:255"`
	PASSWORD string `gorm:"size:255"`
}

type Image struct {
	PATH string `gorm:"size:255"`
	LABEL string `gorm:"size:600"`
}

type Queue struct {
	UserID []User
	ImageID []Image
}
