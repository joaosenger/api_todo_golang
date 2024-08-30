package routes

import (
	"api/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	router.GET("/", services.RouteTest)
	router.GET("/tasks", services.GetAllTasks)
	router.POST("/tasks", services.CreateNewTask)
	router.GET("tasks/:id", services.GetTaskById)
	router.DELETE("tasks/:id", services.DeleteTaskById)
	router.PUT("tasks/:id", services.UpdateTaskById)
}
