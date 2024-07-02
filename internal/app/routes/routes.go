package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/srivalli123125/Todo-Space/internal/app/api"
)

func SetupRoutes(router *gin.Engine, api *api.TodoAPI) {
	v1 := router.Group("/v1")
	{
		v1.POST("/todo", api.CreateTodoHandler())
		v1.PUT("/todo/:id", api.UpdateTodoHandler())
		v1.DELETE("/todo/:id", api.DeleteTodoHandler())
		v1.GET("/todo/:id", api.GetTodoHandler())
		v1.GET("/todo", api.GetAllTodoHandler())
	}
}
