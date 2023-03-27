package apiv1

import (
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

func eventPusherRoute(c *gin.Context) {
	client, err := authenticate(c)
	if err != nil {
		return
	}

	// As long as the client is connected to this route, the client is marked as online
	client.Online = true
	defer func() {
		client.Online = false
	}()

	c.Stream(func(w io.Writer) bool {
		// We don't wait indefinitely on the channels, or gin would not be able to tell if the connection is dead
		timeout := time.NewTimer(1 * time.Second)
		select {
		case peerId := <-client.NewPeers:
			c.SSEvent("new-peer", gin.H{"peerId": peerId})
		case transfer := <-client.NewTransfers:
			c.SSEvent("new-transfer", gin.H{"peerId": transfer.From, "Id": transfer.Id})
		case <-timeout.C:
		}

		// GC the timer to avoid memory issues
		timeout.Stop()

		return true
	})
}
