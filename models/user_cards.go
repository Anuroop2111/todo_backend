package models

type UserCard struct{
	UserID int `gorm:"primaryKey;not null"`
	CardID int `gorm:"primaryKey;not null"`
}