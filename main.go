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

	taskHandler := handlers.NewTaskHandler(db)

	router.POST("/tasks", taskHandler.Create)
	router.GET("/tasks", taskHandler.List)
	router.GET("/tasks/:id", taskHandler.Get)
	router.PUT("/tasks/:id", taskHandler.Update)
	router.DELETE("/tasks/:id", taskHandler.Delete)

	log.Println("Server started on :3003")
	router.Run(":3003")
}