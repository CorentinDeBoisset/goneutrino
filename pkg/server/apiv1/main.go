package apiv1

import (
	"github.com/gin-gonic/gin"
)

func DecorateRouter(r *gin.Engine) error {
	groupV1 := r.Group("/api/v1")

	groupV1.GET("/ping", PingRoute)
	// groupV1.POST("/public-key", PostPublicKeyRoute)

	return nil
}
