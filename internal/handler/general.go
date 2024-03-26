package handler

import (
	"github.com/ali-shokoohi/notes/internal/response"
	"github.com/gin-gonic/gin"
)

type generalHandler struct{}

type GeneralHandler interface {
	Home(c *gin.Context)
}

func NewGeneralHandler() GeneralHandler {
	return &generalHandler{}
}

func (gh *generalHandler) Home(c *gin.Context) {
	response.SendMessage(c, "Hello Bictory!")
}
