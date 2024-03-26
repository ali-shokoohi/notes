package route

import (
	"github.com/ali-shokoohi/notes/internal/handler"
	"github.com/gin-gonic/gin"
)

const BASE = "/api/v1"

func LoadRoutes(engin *gin.Engine, gh handler.GeneralHandler, nh handler.NoteHandler) {
	r := engin.Group(BASE)
	{
		LoadGeneralRoutes(r, gh)
		note := r.Group("/notes")
		{
			LoadNoteRoutes(note, nh)
		}
	}
}
