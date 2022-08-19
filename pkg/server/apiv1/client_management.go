package apiv1

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/corentindeboisset/neutrino/pkg/clientmgr"
	"github.com/corentindeboisset/neutrino/pkg/logger"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/openpgp"
)

// var publicKeyStore map[int]PublicKey

type RegisterBody struct {
	Name      string `json:"name" binding:"required"`
	PublicKey string `json:"publicKey" binding:"required"`
}

func RegisterClientRoute(c *gin.Context) {
	// Receive the public key, the nickname and return a unique numeric id
	// Also set a session cookie, that will be used to re-identify the client
	var body RegisterBody
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	store := c.MustGet("client-store").(*clientmgr.ClientStore)

	// Parse the public key, and re-encode it before saving it to ensure the input is sanitized
	entityList, err := openpgp.ReadArmoredKeyRing(strings.NewReader(body.PublicKey))
	if err != nil {
		logger.ErrorLog("The supplied openPGP public key could not be parsed: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("The supplied public key could not be parsed: %s", err.Error())})
		return
	}
	if len(entityList) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The supplied public key does not contain a valid entity"})
		return
	}

	exp := time.Now().Add(time.Hour * 24 * 30)
	sessionId, err := store.NewClient(body.Name, entityList[0], exp)
	if err != nil {
		logger.ErrorLog("The registration of a new client failed: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occured"})
		return
	}

	c.SetCookie("neutrino-session", sessionId, 24*30*3600, "/", "localhost", true, true)

	c.JSON(200, gin.H{"message": "pong"})
}

func ValidateStatusRoute(c *gin.Context) {
	client, found := c.Get("client")
	if !found {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "There is no current session for the user"})
		return
	}
	clientInstance := client.(*clientmgr.ClientInstance)
	pkeyFingerprint := c.Query("fingerprint")
	if !strings.EqualFold(pkeyFingerprint, clientInstance.PublicKey.PrimaryKey.KeyIdString()) {
		// The supplied public key, and the stored one do not match,
		// even though the client has the right session cookie.
		// Given the inconsistency we return an error and clear everything
		store := c.MustGet("store").(*clientmgr.ClientStore)
		_ = store.RemoveClient(c.GetString("client-uuid"))
		c.SetCookie("neutrino-session", "", -1, "/", "localhost", true, true)

		c.JSON(http.StatusUnauthorized, gin.H{"message": "The submitted public key is invalid"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "The session is valid"})
}

func GetPublicKeyRoute(c *gin.Context) {
	clientId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "the supplied id is not a valid integer"})
		return
	}

	store := c.MustGet("client-store").(*clientmgr.ClientStore)
	client, err := store.GetClientFromId(clientId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "the supplied id could not be associated with another client"})
		return
	}

	if !client.IsOnline() {
		c.JSON(http.StatusNotFound, gin.H{"message": "A client with the given id was found, but is not online"})
		return
	}

	pubKey, err := client.SerializePubKey()
	if err != nil {
		logger.ErrorLog("failed to serialize the public key of a client (id=%d): %s", clientId, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occured"})
		return
	}

	// TODO: send a SSE to the other client if initiate=true

	c.JSON(http.StatusOK, gin.H{"publicKey": pubKey})
}
