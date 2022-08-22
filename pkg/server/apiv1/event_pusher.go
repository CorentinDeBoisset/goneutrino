package apiv1

import (
	"io"
	"net/http"
	"time"

	"github.com/corentindeboisset/neutrino/pkg/clientmgr"
	"github.com/gin-gonic/gin"
)

func eventPusherRoute(c *gin.Context) {
	client, found := c.Get("client")
	if !found {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "A valid session is required to start polling events"})
		return
	}

	clientInstance := client.(clientmgr.ClientInstance)

	// As long as the client is connected to this route, the client is marked as online
	clientInstance.Online = true
	defer func() {
		clientInstance.Online = false
	}()

	c.Stream(func(w io.Writer) bool {
		// We don't wait indefinitely on the channels, or gin would not be able to tell if the connection is dead
		timeout := time.NewTimer(5 * time.Second)
		select {
		case peerId := <-clientInstance.NewPeers:
			c.SSEvent("new-peer", gin.H{"peerId": peerId})
		case transfer := <-clientInstance.NewTransfers:
			c.SSEvent("new-transfer", gin.H{"peerId": transfer.From, "Id": transfer.Id})
		case <-timeout.C:
		}

		// GC the timer to avoid memory issues
		timeout.Stop()

		return true
	})
}
