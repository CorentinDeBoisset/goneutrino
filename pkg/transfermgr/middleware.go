package transfermgr

import "github.com/gin-gonic/gin"

func TransferMgmtMiddleware() gin.HandlerFunc {
	store := InitTransferStore()

	return func(c *gin.Context) {
		// Attach the store to the context
		c.Set("transfer-store", store)
		c.Next()
	}
}
