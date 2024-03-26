package main

import (
	"github.com/ali-shokoohi/notes/internal/handler"
	"github.com/ali-shokoohi/notes/internal/repository"
	"github.com/ali-shokoohi/notes/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	Title string
	Text  string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Note{})

	noteRepository := repository.NewNoteRepository(db)
	noteService := service.NewNoteService(noteRepository)
	noteHandler := handler.NewNoteHandler(noteService)

	r := gin.Default()
	r.GET("/notes", noteHandler.FindAll)
	r.POST("/notes", noteHandler.Create)
	r.Run(":8585") // listen and serve on 0.0.0.0:8080
}
