package apiv1

import (
	"time"

	"github.com/gin-gonic/gin"
)

type PublicKey struct {
	Id         int
	Name       string
	PublicKey  []byte
	Expiration time.Time
}

// var publicKeyStore map[int]PublicKey

func GenerateIdentifierRoute(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

func PostPublicKeyRoute(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

func UpdatePublicKeyRoute(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}
