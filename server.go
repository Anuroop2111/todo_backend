package main

import (
	"github.com/Anuroop2111/todo_backend/database"
)

func init() {
	database.ConnectDB()
}

func main() {
	// database.DB.AutoMigrate(&models.User{})
}

