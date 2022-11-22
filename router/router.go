package router

import (
	handlers "transactions/app/handlers"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/health", handlers.Health)
		api.POST("/transactions", handlers.Add)
		api.DELETE("/transactions", handlers.Destroy)
		api.PUT("/setloc/:id", handlers.SetLocation)
		api.PUT("/resetloc/:id", handlers.ResetLocation)
		api.GET("/statistics", handlers.Stats)
	}
	return r
}
