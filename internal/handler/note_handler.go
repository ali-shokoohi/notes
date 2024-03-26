package handler

import (
	"net/http"

	"github.com/ali-shokoohi/notes/internal/service"
	"github.com/gin-gonic/gin"
)

type noteHandler struct {
	noteService service.NoteService
}

type NoteHandler interface {
	Create(c *gin.Context)
	FindAll(c *gin.Context)
}

func NewNoteHandler(
	ns service.NoteService,
) NoteHandler {
	return &noteHandler{
		noteService: ns,
	}
}

func (nh *noteHandler) Create(c *gin.Context) {}

func (nh *noteHandler) FindAll(c *gin.Context) {
	notes, err := nh.noteService.FindAll()
	if err != nil {
	}

	c.JSON(http.StatusOK, notes)
}
