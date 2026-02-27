package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET ALL OPENING",
	})
}
