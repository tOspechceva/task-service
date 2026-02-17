package main

import (
	"log"

	"task-service/config"
	"task-service/database"
	"task-service/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env")
	}

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}

	defer db.Close()

	// =========================
	// МИГРАЦИИ
	// =========================
	err = database.RunMigrations(db)
	if err != nil {
		log.Fatal("Ошибка миграции:", err)
	}

	// =========================
	// SEED DATA
	// =========================
	err = database.SeedData(db)
	if err != nil {
		log.Fatal("Ошибка seed:", err)
	}

	log.Println("База данных готова")

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	router.GET("/tasks", handlers.GetTasks)

	log.Println("Server started on :3003")
	router.Run(":3003")
}