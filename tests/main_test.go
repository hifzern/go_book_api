package tests

import (
	"bytes"
	"encoding/json"
	"go_book_api/api"
	"net/http"
	"net/http/httptest"
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
	router.POST("/book", api.CreateBook)

	book := api.Book{
		Title:  "Raja Gorong-gorong",
		Author: "Girban",
		Year:   2023,
	}
	jsonValue, _ := json.Marshal(book)
	req, _ := http.NewRequest("POST", "/book", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusCreated {
		t.Errorf("Expected status %d got %d", http.StatusCreated, status)
	}
	var response api.JsonResponse
	json.NewDecoder(w.Body).Decode(&response)

	if response.Data == nil {
		t.Errorf("Expected book data got nil")
	}
}
