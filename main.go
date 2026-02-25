// main.go
package main

import (
	"database/sql"
	"log"
	"net/http"


	"github.com/gin-gonic/gin"
	
	_ "github.com/lib/pq" // или другой драйвер БД
	
	"task-service/handlers"
	"task-service/repository"
	"task-service/services"
)

func main() {
	// =====================
	// 1. ПОДКЛЮЧЕНИЕ К БД
	// =====================
	dsn := "postgres://postgres:qwerty123@localhost:5432/task_db?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("❌ Failed to connect to DB: %v", err)
	}
	
	if err := db.Ping(); err != nil {
		log.Fatalf("❌ DB ping failed: %v", err)
	}
	defer db.Close()
	
	log.Println("✅ Database connected")

	// =====================
	// 2. СОЗДАНИЕ ЗАВИСИМОСТЕЙ (Dependency Injection)
	// =====================
	
	// DB → Repository
	taskRepo := repository.NewTaskRepository(db)
	
	// Repository → Service
	taskService := services.NewTaskService(taskRepo)
	
	// Service → Handler
	taskHandler := handlers.NewTaskHandler(taskService) // 👈 Теперь передаём service, а не db

	// =====================
	// 3. НАСТРОЙКА GIN
	// =====================
	r := gin.Default()
	


	// =====================
	// 4. РЕГИСТРАЦИЯ РОУТОВ
	// =====================
	tasks := r.Group("/tasks")
	{
		tasks.GET("", taskHandler.List)           // GET /tasks
		tasks.GET("/:id", taskHandler.Get)        // GET /tasks/:id
		tasks.POST("", taskHandler.Create)        // POST /tasks
		tasks.PUT("/:id", taskHandler.Update)     // PUT /tasks/:id
		tasks.DELETE("/:id", taskHandler.Delete)  // DELETE /tasks/:id
	}

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// =====================
	// 5. ЗАПУСК СЕРВЕРА
	// =====================
	log.Println("🚀 Server starting on :3003")
	if err := r.Run(":3003"); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}