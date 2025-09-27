package api

import "github.com/gin-gonic/gin"

type Book struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	Year	int    `json:"year" binding:"required,gte=1000,lte=2100"`
	CoverURL string `json:"cover_url"`
	Description string `json:"description"`
	Genre     string `json:"genre"`

}

type JsonResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func ResponseJSON(c *gin.Context, status int, message string, data any) {
	resp := gin.H{
		"status":  status,
		"message": message,
	}
	if data != nil {
		resp["data"] = data
	}
	c.JSON(status, resp)
}



