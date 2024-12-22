package database

import (
	"tasksmanager/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	DB, err = gorm.Open("postgres", "host=localhost user=postgres dbname=taskmanager password=root sslmode=disable")
	if err != nil {
		panic("failed to connect to database")
	}

	DB.AutoMigrate(&models.Task{}, &models.Entry{})
}
