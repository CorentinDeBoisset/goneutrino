package apiv1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func pingRoute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "api__pong"})
}
