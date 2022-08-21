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

type RegisterBody struct {
	Name      string `json:"name" binding:"required"`
	PublicKey string `json:"publicKey" binding:"required"`
}

func registerClientRoute(c *gin.Context) {
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

	// TODO: send the expiration date of the session to the JS
	c.JSON(200, gin.H{"message": "Ok"})
}

func validateStatusRoute(c *gin.Context) {
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

	// TODO: send the expiration of the session to the JS
	c.JSON(http.StatusOK, gin.H{"message": "The session is valid"})
}

func getPublicKeyRoute(c *gin.Context) {
	rawClient, found := c.Get("client")
	if !found {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	client := rawClient.(clientmgr.ClientInstance)

	clientId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "the supplied value for `id` is not a valid integer"})
		return
	}
	initiate, err := strconv.ParseBool(c.DefaultQuery("initiate", "false"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "the supplied value for `initiate` is not a valid boolean"})
	}

	store := c.MustGet("client-store").(*clientmgr.ClientStore)
	peerClient, err := store.GetClientFromId(clientId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "the supplied id could not be associated with another client"})
		return
	}

	if !peerClient.IsOnline() {
		c.JSON(http.StatusNotFound, gin.H{"message": "A client with the given id was found, but is not online"})
		return
	}

	pubKey, err := peerClient.SerializePubKey()
	if err != nil {
		logger.ErrorLog("failed to serialize the public key of a client (id=%d): %s", clientId, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "an error occured"})
		return
	}

	if initiate {
		// This call may be blocking, not sure if it should be put in a specific goroutine...
		// In that case, there may be too many goroutines started
		// Alternatively, we could use a buffered channel and a non-blocking send, and handle the error if many other clients are trying to connect to the peer
		peerClient.NewPeers <- client.Id
	}

	c.JSON(http.StatusOK, gin.H{"publicKey": pubKey})
}
