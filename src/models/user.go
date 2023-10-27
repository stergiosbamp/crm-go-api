package models

type User struct {
	ID       uint   
	Username string `gorm:"NOT NULL;index:,unique"`
	Password string `gorm:"NOT NULL"`
}
