package service

import (
	"log"

	"github.com/ali-shokoohi/notes/internal/model"
	"github.com/ali-shokoohi/notes/internal/repository"
)

type noteService struct {
	noteRepository repository.NoteRepository
}

type NoteService interface {
	Create(title, text string) error
	FindAll() ([]*model.Note, error)
}

func NewNoteService(noteRepository repository.NoteRepository) NoteService {
	return &noteService{
		noteRepository: noteRepository,
	}
}

func (ns *noteService) Create(title, text string) error {
	err := ns.noteRepository.Create(title, text)
	if err != nil {
		log.Println("Failure at create note in service layer: ", err)
		return err
	}
	return nil
}

func (ns *noteService) FindAll() ([]*model.Note, error) {
	notes, err := ns.noteRepository.FindAll()
	return notes, err
}
