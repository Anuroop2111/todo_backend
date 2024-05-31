package models

type CardItem struct {
	CardID int `gorm:"primaryKey;not null"`
	UserID int `gorm:"primaryKey;not null"`
}