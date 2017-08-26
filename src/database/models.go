package database

type User struct {
	ID       int     `gorm:"primary_key; AUTO_INCREMENT"`
	LOGIN    string  `gorm:"size:255; unique; not null"`
	PASSWORD string  `gorm:"size:255; not null"`
	UUID     string  `gorm:"size:255"`
	Image    []Image `gorm:"ForeignKey:ImageID"`
	ImageID  uint
}

type Image struct {
	ID       int    `gorm:"primary_key"`
	FILENAME string `gorm:"size:255"`
	User     User   `gorm:"ForeignKey:UserID"`
	UserID   uint
}

type Queue struct {
	ID      int   `gorm:"primary_key"`
	User    User  `gorm:"ForeignKey:UserID"`
	Image   Image `gorm:"ForeignKey:ImageID"`
	UserID  uint
	ImageID uint
}
