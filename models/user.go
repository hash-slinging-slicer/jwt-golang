package models

type User struct {
	Id       int64  `gorm:"primaryKey" json="id"`
	Nama     string `gorm:"varchar(255)" json:"nama"`
	Username string `gorm:"varchar(255)" json:"username"`
	Password string `gorm:"varchar(255)" json:"password"`
}
