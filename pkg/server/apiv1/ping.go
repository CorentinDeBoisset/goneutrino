package apiv1

import "github.com/gin-gonic/gin"

func pingRoute(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}
