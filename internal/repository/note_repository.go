package repository

import (
	"log"

	"github.com/ali-shokoohi/notes/internal/model"
	"gorm.io/gorm"
)

type noteRepository struct {
	db *gorm.DB
}

type NoteRepository interface {
	Create(title, text string) error
	FindAll() ([]*model.Note, error)
}

func NewNoteRepository(
	db *gorm.DB,
) NoteRepository {
	return &noteRepository{
		db: db,
	}
}

func (nr *noteRepository) Create(title, text string) error {
	note := model.Note{
		Title: title,
		Text:  text,
	}
	err := nr.db.Create(&note).Error
	if err != nil {
		log.Printf("A failure at creating a new Note: %+v\n", err)
		return err
	}
	return nil
}

func (nr *noteRepository) FindAll() ([]*model.Note, error) {
	var notes []*model.Note
	err := nr.db.Find(&notes).Error
	return notes, err
}
