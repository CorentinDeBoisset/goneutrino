package apiv1

import "github.com/gin-gonic/gin"

func SendStringRoute(c *gin.Context) {
	// Store the sent encrypted string in memory, with a unique id
	// Send an SSE to the recipient, with the unique id
	c.JSON(200, gin.H{"message": "pong"})
}

func GetStringRoute(c *gin.Context) {
	// Use the given unique id
	// Send the client the encrypted string
	c.JSON(200, gin.H{"message": "pong"})
}

func SendFileRoute(c *gin.Context) {
	// Same as for strings, but using a buffer between the sender and the recipient
	c.JSON(200, gin.H{"message": "pong"})
}

func GetFileRoute(c *gin.Context) {
	// Same as for strings, but using a buffer between the sender and the recipient
	c.JSON(200, gin.H{"message": "pong"})
}
