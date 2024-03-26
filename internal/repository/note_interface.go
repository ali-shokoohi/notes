package repository

import "github.com/ali-shokoohi/notes/internal/model"

type NoteRepository interface {
	CommonBehaviourRepository[model.Note]
}
