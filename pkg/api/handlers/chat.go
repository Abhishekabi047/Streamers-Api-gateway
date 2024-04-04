package handlers

import (
	"api/pkg/client/interfaces"
	"api/pkg/models"
	"api/pkg/utils"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ChatHandler struct {
	Client interfaces.ChatClient
}

func NewChatHandler(client interfaces.ChatClient) ChatHandler {
	return ChatHandler{
		Client: client,
	}
}

func (a *ChatHandler) ChatHistory(c *gin.Context) {
	u1 := c.Query("u1")
	u2 := c.Query("u2")

	fromTS, toTS := "0", "+inf"
	if c.Query("from-ts") != "" && c.Query("to-ts") != "" {
		fromTS = c.Query("from-ts")
		toTS = c.Query("to-ts")
	}
	body := models.ChathistoryRequest{
		Username1: u1,
		Username2: u2,
		Fromts:    fromTS,
		Tots:      toTS,
	}
	res, err := a.Client.ChatHistory(context.Background(), body)
	if err != nil {
		errmsg := utils.ExtractError(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errmsg,
			"error":   err,
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (a *ChatHandler) ContactList(c *gin.Context) {
	u := c.Query("username")

	res, err := a.Client.ContactList(context.Background(), u)
	if err != nil {
		errmsg := utils.ExtractError(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errmsg,
			"error":   err,
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (a *ChatHandler) VerifyContact(c *gin.Context) {
	u := c.Query("username")
	res, err := a.Client.VerifyContact(context.Background(), u)
	if err != nil {
		errmsg := utils.ExtractError(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errmsg,
			"error":   err,
		})
		return
	}
	c.JSON(http.StatusOK, res)
}
