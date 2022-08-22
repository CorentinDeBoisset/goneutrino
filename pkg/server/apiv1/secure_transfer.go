package apiv1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/corentindeboisset/neutrino/pkg/clientmgr"
	"github.com/corentindeboisset/neutrino/pkg/logger"
	"github.com/corentindeboisset/neutrino/pkg/transfermgr"
	"github.com/gin-gonic/gin"
)

func sendStringRoute(c *gin.Context) {
	rawCurrentClient, found := c.Get("client")
	if !found {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "An ongoing session is required to upload files"})
		return
	}
	curClient := rawCurrentClient.(*clientmgr.ClientInstance)

	destId, err := strconv.Atoi(c.Query("destination-id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "A valid destination ID is required"})
		return
	}

	secret := c.Query("secret")
	if len(secret) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No valid file could be found in the upload"})
		return
	}

	clientStore := c.MustGet("client-store").(*clientmgr.ClientStore)
	destClient, err := clientStore.GetClientFromId(destId)
	if err != nil || !destClient.IsOnline() {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No client is associated to the submitted id"})
		return
	}

	transferStore := c.MustGet("transfer-store").(*transfermgr.TransferStore)
	newTransfer, err := transferStore.NewStringTransfer(curClient.Id, destClient.Id, secret)
	if err != nil {
		logger.ErrorLog("An error occured when creating and saving the transfer: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An error occured"})
		return
	}

	// Same problem as the route getPublicKeyRoute() - in client_management.go
	destClient.NewTransfers <- newTransfer

	c.JSON(http.StatusNoContent, nil)
}

func getStringRoute(c *gin.Context) {
	transferStore := c.MustGet("transfer-store").(*transfermgr.TransferStore)
	rawCurrentClient, found := c.Get("client")
	if !found {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "An ongoing session is required to upload files"})
		return
	}
	curClient := rawCurrentClient.(*clientmgr.ClientInstance)

	transferId := c.Query("transfer-id")
	transfer, err := transferStore.GetTransfer(transferId)
	if err != nil || transfer.To != curClient.Id {
		c.JSON(http.StatusNotFound, gin.H{"message": "There are no transfer with this identifier"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"secret": transfer.TransferedString})
}

func sendFileRoute(c *gin.Context) {
	rawCurrentClient, found := c.Get("client")
	if !found {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "An ongoing session is required to upload files"})
		return
	}
	curClient := rawCurrentClient.(*clientmgr.ClientInstance)

	destId, err := strconv.Atoi(c.Query("destination-id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "A valid destination ID is required"})
		return
	}

	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No valid file could be found in the upload"})
		return
	}

	clientStore := c.MustGet("client-store").(*clientmgr.ClientStore)
	destClient, err := clientStore.GetClientFromId(destId)
	if err != nil || !destClient.IsOnline() {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No client is associated to the submitted id"})
		return
	}

	transferStore := c.MustGet("transfer-store").(*transfermgr.TransferStore)
	newTransfer, err := transferStore.NewFileTransfer(curClient.Id, destClient.Id, fileHeader)
	if err != nil {
		logger.ErrorLog("An error occured when creating and saving the transfer: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An error occured"})
		return
	}

	// Same problem as the route getPublicKeyRoute() - in client_management.go
	destClient.NewTransfers <- newTransfer

	c.JSON(http.StatusNoContent, nil)
}

func getFileRoute(c *gin.Context) {
	transferStore := c.MustGet("transfer-store").(*transfermgr.TransferStore)
	rawCurrentClient, found := c.Get("client")
	if !found {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "An ongoing session is required to upload files"})
		return
	}
	curClient := rawCurrentClient.(*clientmgr.ClientInstance)

	transferId := c.Query("transfer-id")
	transfer, err := transferStore.GetTransfer(transferId)
	if err != nil || transfer.To != curClient.Id {
		c.JSON(http.StatusNotFound, gin.H{"message": "There are no transfer with this identifier"})
		return
	}

	file, err := transfer.TransferedFile.Open()
	if err != nil {
		logger.ErrorLog("An error occured when trying to access the file from a transfer: %s", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An error occured"})
		return
	}

	extraHeaders := map[string]string{
		"Content-Disposition": fmt.Sprintf("attachment; filename=\"%s\"", transfer.TransferedFile.Filename),
	}

	c.DataFromReader(http.StatusOK, transfer.TransferedFile.Size, "content: bytes", file, extraHeaders)
}
