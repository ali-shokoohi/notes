package route

import (
	"github.com/ali-shokoohi/notes/internal/handler"
	"github.com/gin-gonic/gin"
)

func LoadNoteRoutes(group *gin.RouterGroup, nh handler.NoteHandler) {
	group.POST("", nh.CreateNote)
	group.GET("", nh.GetNotesWithPagination)
	group.GET("/filter", nh.GetNotesByFilterWithPagination)
	group.GET("/:id", nh.GetNoteByID)
	group.PUT("/:id", nh.UpdateNote)
	group.DELETE("/:id", nh.DeleteNote)
}
