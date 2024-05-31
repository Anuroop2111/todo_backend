package models

import (
	"time"
)

type TagType string

const (
	Favourite TagType = "fav"
	Important TagType = "imp"
)

type Card struct {
	CardUserID int `gorm:"primaryKey;autoIncrement;not null"`
	CardName string `gorm:"not null;size:45"`
	Tag        TagType   `gorm:"type:enum('fav', 'imp');"`
	StartTime  time.Time `gorm:"not null"`
	EndTime    *time.Time
}