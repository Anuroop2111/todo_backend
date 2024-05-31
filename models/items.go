package models

type Item struct {
	ItemID          int    `gorm:"primaryKey;autoIncrement;not null"`
	ItemName        string `gorm:"not null;size:45"`
	ItemDescription string `gorm:"type:text"`
	FinishStatus    bool   `gorm:"default:false"`
}