package apiv1

import "github.com/gin-gonic/gin"

func sendStringRoute(c *gin.Context) {
	// We can store the string in memory, it’s ok
	c.JSON(200, gin.H{"message": "pong"})
}

func getStringRoute(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

func sendFileRoute(c *gin.Context) {
	// Find a way to send the file reader to the getFileRoute from the other client, and pipeline the data using c.DataFromReader()
	c.JSON(200, gin.H{"message": "pong"})
}

func getFileRoute(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}
