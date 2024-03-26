package gorm

import (
	"github.com/ali-shokoohi/notes/internal/model"
	"github.com/ali-shokoohi/notes/internal/repository"
	"gorm.io/gorm"
)

type noteRepository struct {
	repository.CommonBehaviourRepository[model.Note]
}

func NewNoteRepository(db *gorm.DB) repository.NoteRepository {
	return &noteRepository{
		CommonBehaviourRepository: NewCommonBehaviour[model.Note](db),
	}
}
