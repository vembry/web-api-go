package handler

import "github.com/gin-gonic/gin"

func InitHandler(r *gin.Engine) {

	healthGroup := r.Group("/health")
	{
		healthGroup.GET("", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
}
