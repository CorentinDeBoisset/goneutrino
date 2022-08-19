package apiv1

import (
	"io"

	"github.com/gin-gonic/gin"
)

func EventPusherRoute(c *gin.Context) {
	c.Stream(func(w io.Writer) bool {
		// Send SSE with c.SSEvent("messageType", gin.H{"message": "pong"})
		return true
	})
}
