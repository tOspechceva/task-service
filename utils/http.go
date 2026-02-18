package utils

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Error(c *gin.Context, status int, err error) {

	// лог в консоль сервера
	log.Println("ERROR:", err)

	c.JSON(status, gin.H{
		"success": false,
		"error": gin.H{
			"message": err.Error(),
		},
	})
}
