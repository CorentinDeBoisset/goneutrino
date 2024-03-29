package apiv1

import (
	"net/http"
	"strings"
	"time"

	"github.com/ProtonMail/go-crypto/openpgp"
	"github.com/corentindeboisset/neutrino/pkg/clientmgr"
	"github.com/corentindeboisset/neutrino/pkg/logger"
	"github.com/gin-gonic/gin"
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
		c.JSON(http.StatusBadRequest, gin.H{"msg": "api__invalid_format", "error": err.Error()})
		return
	}

	store := c.MustGet("client-store").(*clientmgr.ClientStore)

	// Parse the public key, and re-encode it before saving it to ensure the input is sanitized
	entityList, err := openpgp.ReadArmoredKeyRing(strings.NewReader(body.PublicKey))
	if err != nil {
		logger.ErrorLog("The supplied openPGP public key could not be parsed: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"msg": "api__invalid_pgpkey"})
		return
	}
	if len(entityList) == 0 {
		logger.ErrorLog("The supplied openPGP public key did not contain any entity")
		c.JSON(http.StatusBadRequest, gin.H{"msg": "api__invalid_pgpkey"})
		return
	}

	exp := time.Now().Add(time.Hour * 24 * 30)
	sessionId, err := store.NewClient(body.Name, entityList[0], exp)
	if err != nil {
		logger.ErrorLog("The registration of a new client failed: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "api__generic_error"})
		return
	}

	// TODO: make the cookie secure in production mode
	c.SetCookie("neutrino-session", sessionId, 24*30*3600, "/", "", false, true)

	// This is a js-readable cookie to check if there is an existing session
	c.SetCookie("neutrino-js-session", "", 24*30*3600, "/", "", false, false)

	// Check the client
	client, err := store.GetClient(sessionId)
	if err != nil {
		// This should never happen...
		logger.ErrorLog("Failed to find the client just after it was registered")
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "api__generic_error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "api__ok", "payload": gin.H{"id": client.Id}})
}

// FIXME: remove this route
// func getPublicKeyRoute(c *gin.Context) {
// 	rawClient, found := c.Get("client")
// 	if !found {
// 		c.JSON(http.StatusUnauthorized, gin.H{"msg": "api__unauthorized"})
// 		return
// 	}
// 	client := rawClient.(clientmgr.ClientInstance)

// 	clientId, err := strconv.Atoi(c.Query("id"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"msg": "api__invalid_format", "error": "The supplied value for `id` is not an integer"})
// 		return
// 	}
// 	initiate, err := strconv.ParseBool(c.DefaultQuery("initiate", "false"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"msg": "api__invalid_format", "error": "the supplied value for `initiate` is not a valid boolean"})
// 	}

// 	store := c.MustGet("client-store").(*clientmgr.ClientStore)
// 	peerClient, err := store.GetClientFromId(clientId)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"msg": "api__unknown_peer", "error": "the supplied id could not be associated with another client"})
// 		return
// 	}

// 	if !peerClient.IsOnline() {
// 		c.JSON(http.StatusNotFound, gin.H{"msg": "api__peer_offline", "error": "A client with the given id was found, but is not online"})
// 		return
// 	}

// 	pubKey, err := peerClient.SerializePubKey()
// 	if err != nil {
// 		logger.ErrorLog("failed to serialize the public key of a client (id=%d): %s", clientId, err.Error())
// 		c.JSON(http.StatusInternalServerError, gin.H{"msg": "api__generic_error"})
// 		return
// 	}

// 	if initiate {
// 		// This call may be blocking, not sure if it should be put in a specific goroutine...
// 		// In that case, there may be too many goroutines started
// 		// Alternatively, we could use a buffered channel and a non-blocking send, and handle the error if many other clients are trying to connect to the peer
// 		peerClient.NewPeers <- client.Id
// 	}

// 	c.JSON(http.StatusOK, gin.H{"msg": "api__ok", "publicKey": pubKey})
// }
