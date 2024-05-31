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
	CardID int `gorm:"primaryKey;autoIncrement;not null"`
	CardName string `gorm:"not null;size:45"`
	Tag        TagType   `gorm:"type:enum('fav', 'imp');"`
	StartTime  time.Time `gorm:"not null"`
	EndTime    *time.Time
	Users      []User `gorm:"many2many:user_cards;joinForeignKey:CardID;joinReferences:UserID"`
	Items []Item `gorm:"many2many:card_items;joinForeignKey:CardID;joinReferences:ItemID"`
}
