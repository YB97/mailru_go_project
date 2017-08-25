package database

type User struct {
	ID       int    `gorm:"primary_key"`
	LOGIN    string `gorm:"size:255"`
	PASSWORD string `gorm:"size:255"`
	UUID     string `gorm:"size:255"`
}

type Image struct {
	ID       int    `gorm:"primary_key"`
	FILENAME string `gorm:"size:255"`
	LABEL    string `gorm:"size:600"`
}

type Queue struct {
	ID      int     `gorm:"primary_key"`
	User User `gorm:"ForeignKey:UserID"`
	Image Image `gorm:"ForeignKey:ImageID"`
	UserID uint
	ImageID uint
}
