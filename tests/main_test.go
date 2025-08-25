package tests

import (
	"go_book_api/api"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() {
	var err error
	api.DB, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	api.DB.AutoMigrate(&api.Book{})
}

func addBook() api.Book {
	book := api.Book{Title: "Raja Solo", Author: "Wowok Gendut", Year: 2023}
	api.DB.Create(&book)
	return book
}

func TestCreateBook(t *testing.T) {
	setupTestDB()
	router := gin.Default()
}
