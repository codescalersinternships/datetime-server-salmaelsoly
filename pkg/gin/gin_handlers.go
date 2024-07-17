package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func SetUpRouter() *gin.Engine {
	server := gin.Default()
	server.GET("/datetime", DateTime())
	return server
}

func DateTime() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": time.Now().Format(time.UnixDate),
		})
	}
}
