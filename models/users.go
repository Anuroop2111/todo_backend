package models

type User struct {
	UserID   int    `gorm:"primaryKey;autoIncrement;not null"`
	UserName string `gorm:"not null;size:45"`
	Email    string `gorm:"not null;size:45"`
	Password string `gorm:"not null;size:100"`
}