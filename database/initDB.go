package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/Anuroop2111/todo_backend/models"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	DB, err = gorm.Open(mysql.Open("root:1234@/todo_db"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to mysql database")
	}

	err = DB.AutoMigrate(&models.User{}, &models.Card{}, &models.Item{}, &models.UserCard{})
	if err != nil {
		panic("Failed to auto-migrate user table")
	}
}