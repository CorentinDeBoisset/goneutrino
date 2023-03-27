package apiv1

import (
	"fmt"
	"net/http"

	"github.com/corentindeboisset/neutrino/pkg/clientmgr"
	"github.com/gin-gonic/gin"
)

func authenticate(c *gin.Context) (*clientmgr.ClientInstance, error) {
	rawClient, found := c.Get("client")
	if !found {
		// Return a 403, and reset all cookies
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "api__unauthorized"})
		c.SetCookie("neutrino-session", "", -1, "/", "", false, true)
		c.SetCookie("neutrino-js-session", "", -1, "/", "", false, false)

		return nil, fmt.Errorf("failed to authenticate the client")
	}
	client := rawClient.(clientmgr.ClientInstance)

	return &client, nil
}
