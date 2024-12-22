package router

import (
	"tasksmanager/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Настройка CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Разрешить запросы с любых источников
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true, // Разрешить отправку cookie
		ExposeHeaders:    []string{"Content-Length"},
	}))

	taskRoutes := router.Group("/tasks")
	{
		taskRoutes.POST("/", handlers.CreateTask)
		taskRoutes.GET("/", handlers.GetTasks)
		taskRoutes.PUT("/:id", handlers.UpdateTask)
		taskRoutes.DELETE("/:id", handlers.DeleteTask)
	}

	entryRoutes := router.Group("/entries")
	{
		entryRoutes.POST("/", handlers.CreateEntry)      // Создание записи
		entryRoutes.GET("/", handlers.GetEntries)        // Получение всех записей
		entryRoutes.GET("/:id", handlers.GetEntryByID)   // Получение записи по ID
		entryRoutes.PUT("/:id", handlers.UpdateEntry)    // Обновление записи
		entryRoutes.DELETE("/:id", handlers.DeleteEntry) // Удаление записи
	}

	return router
}
