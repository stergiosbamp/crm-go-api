package models

import "time"

type Token struct {
	ID        uint
	UserID    uint      `gorm:"NOT NULL"`
	User      *User     `gorm:"NOT NULL"`
	Token     string    `gorm:"NOT NULL"`
	Status    string    `gorm:"NOT NULL;type:enum('active','inactive')"`
	CreatedAt time.Time `gorm:"NOT NULL;autoCreateTime"`
}
