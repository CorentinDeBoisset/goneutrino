package apiv1

import (
	"github.com/gin-gonic/gin"
)

func DecorateRouter(r *gin.Engine) {
	groupV1 := r.Group("/api/v1")

	groupV1.GET("/ping", pingRoute)
	groupV1.POST("/register", registerClientRoute)
	groupV1.GET("/validate-status", validateStatusRoute)
	groupV1.GET("/public-key", getPublicKeyRoute)

	groupV1.GET("/SSE", eventPusherRoute)

	groupV1.POST("/secret/string", sendStringRoute)
	groupV1.GET("/secret/string", getStringRoute)
	groupV1.POST("/secret/file", sendFileRoute)
	groupV1.GET("/secret/file", getFileRoute)
}
