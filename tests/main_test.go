package tests

import (
	"bytes"
	"encoding/json"
	"go_book_api/api"
	"net/http"
	"net/http/httptest"
	"strconv"
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

func TestGetBooks(t *testing.T) {
	setupTestDB()
	addBook()
	router := gin.Default()
	router.GET("/books", api.GetBooks)

	req, _ := http.NewRequest("GET", "/books", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("expected status %d got %d", http.StatusOK, status)
	}

	var response api.JsonResponse
	json.NewDecoder(w.Body).Decode(&response)

	if len(response.Data.([]interface{})) == 0 {
		t.Errorf("Expected non empty book list")
	}
}

func TestUpdateBook(t *testing.T) {
	setupTestDB()
	book := addBook()
	router := gin.Default()
	router.PUT("/book/:id", api.UpdateBook)

	updateBook := api.Book{
		Title: "", Author: "", Year: 2024,
	}
	jsonValue, _ := json.Marshal(updateBook)
	req, _ := http.NewRequest("PUT", "/book/"+strconv.Itoa(int(book.ID)), bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("expected status %d got %d", http.StatusOK, status)
	}

	var response api.JsonResponse
	json.NewDecoder(w.Body).Decode(&response)

	if response.Data == nil || response.Data.(map[string]interface{})["title"] != "Raja Solo" {
		t.Errorf("Expected update book title got %v", response.Data)
	}
}
