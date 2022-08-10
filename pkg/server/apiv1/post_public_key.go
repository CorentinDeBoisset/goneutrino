package apiv1

import "github.com/gin-gonic/gin"

func PostPublicKeyRoute(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}
