package router

import (
	"gopportunities/handler"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1");

	{
		v1.GET("/openings", handler.GetAllOpeningHandler)
		v1.POST("/openings", handler.CreateOpeningHandler)
		v1.POST("/openings/:id", handler.GetOpeningHandler)
		v1.PUT("/openings/:id", handler.UpdateOpeningHandler)
		v1.DELETE("/openings/:id", handler.DeleteOpeningHandler)
	}

}