package handlers

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {

	tasks := []map[string]string{
		{"id": "1", "title": "Первая задача"},
		{"id": "2", "title": "Вторая задача"},
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    tasks,
	})
}