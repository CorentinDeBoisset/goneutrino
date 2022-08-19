package server

import (
	"time"

	"github.com/corentindeboisset/neutrino/pkg/clientmgr"
	"github.com/gin-gonic/gin"
)

func clientMgmtMiddleware() gin.HandlerFunc {
	store := clientmgr.InitStore()

	// Start a task that cleans up old sessions every 15 minutes
	go func() {
		ticker := time.NewTicker(time.Minute * 15)
		for range ticker.C {
			store.CleanupExpired()
		}
	}()

	return func(c *gin.Context) {
		// Attach the store to the context
		c.Set("client-store", store)

		sessionId, err := c.Cookie("neutrino-session")
		if err != nil {
			// There is no session cookie, we can stop
			c.Next()
			return
		}

		client, err := store.GetClient(sessionId)
		if err != nil {
			c.SetCookie("neutrino-session", "", -1, "/", "localhost", true, true)
			c.Next()
			return
		}

		c.Set("client-uuid", sessionId)
		c.Set("client", client)

		c.Next()
	}
}
