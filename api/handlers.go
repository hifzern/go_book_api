package api

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	_ = godotenv.Load()
	dsn := os.Getenv("DB_URL")

	if dsn == "" {
		log.Fatal("DB_URL is empty")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})

	if err != nil {
		log.Fatal("Failed to connect to database : ", err)
	}

	//test connection
	sqlDB, err := db.DB()

	if err != nil {
		log.Fatal("Failed to get sql DB from gorm : ", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatal("Failed to ping database : ", err)
	}
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)

	DB = db

	//migrate the schema
	if err := DB.AutoMigrate(&Book{}); err != nil {
		log.Fatal("Failed to migrate schema : ", err)
	}
}

func CreateBook(c *gin.Context) {
	var book Book

	//bind the request body
	if err := c.ShouldBindJSON(&book); err != nil {
		ResponseJSON(c, http.StatusBadRequest, "Invalid input", nil)
		return
	}
	if err := DB.WithContext(c.Request.Context()).Create(&book).Error; err != nil {
		ResponseJSON(c, http.StatusInternalServerError, "Failed to create book", err.Error())
		return
	}
	ResponseJSON(c, http.StatusCreated, "Book created successfully", book)
}

func GetBooks(c *gin.Context) {
	var books []Book

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	offset := (page - 1) * limit

	if err := DB.Limit(limit).Offset(offset).Find(&books).Error; err != nil {
		ResponseJSON(c, http.StatusInternalServerError, "Failed to retrieve books", nil)
		return
	}

	ResponseJSON(c, http.StatusOK, "Books retrieved successfully", books)
}


func GetBook(c *gin.Context) {
    var book Book
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        ResponseJSON(c, http.StatusBadRequest, "Invalid ID", nil)
        return
    }

    if err := DB.First(&book, id).Error; err != nil {
        ResponseJSON(c, http.StatusNotFound, "Book not found", nil)
        return
    }

    ResponseJSON(c, http.StatusOK, "Book retrieved successfully", book)
}


func UpdateBook(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        ResponseJSON(c, http.StatusBadRequest, "Invalid ID", nil)
        return
    }

    var book Book
    if err := DB.First(&book, id).Error; err != nil {
        ResponseJSON(c, http.StatusNotFound, "Book not found", nil)
        return
    }

    var input Book
    if err := c.ShouldBindJSON(&input); err != nil {
        ResponseJSON(c, http.StatusBadRequest, "Invalid input", err.Error())
        return
    }

    if err := DB.Model(&book).Updates(input).Error; err != nil {
        ResponseJSON(c, http.StatusInternalServerError, "Failed to update book", nil)
        return
    }

    ResponseJSON(c, http.StatusOK, "Book updated successfully", book)
}


func DeleteBook(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        ResponseJSON(c, http.StatusBadRequest, "Invalid ID", nil)
        return
    }

    result := DB.Delete(&Book{}, id)
    if result.RowsAffected == 0 {
        ResponseJSON(c, http.StatusNotFound, "Book not found", nil)
        return
    }

    ResponseJSON(c, http.StatusOK, "Book deleted successfully", nil)
}

