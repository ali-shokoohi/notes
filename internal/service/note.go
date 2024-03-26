package service

import (
	"context"
	"errors"
	"time"

	"github.com/ali-shokoohi/notes/internal/config"
	"github.com/ali-shokoohi/notes/internal/dto"
	customErrors "github.com/ali-shokoohi/notes/internal/errors"
	"github.com/ali-shokoohi/notes/internal/model"
	"github.com/ali-shokoohi/notes/internal/repository"
	"go.uber.org/zap"
)

// noteService implements the NoteService interface.
// It holds the configuration settings and the repository required to manage notes.
type noteService struct {
	cfg            *config.Config            // cfg contains the application's configurations.
	noteRepository repository.NoteRepository // noteRepository is the interface for the Note repository.
}

// NoteService defines the operations available for managing notes.
type NoteService interface {
	CreateNote(ctx context.Context, noteDTO dto.CreateNoteRequest) error                                                                              // CreateNote creates a new note based on the provided DTO.
	UpdateNote(ctx context.Context, id uint, noteDTO dto.UpdateNoteRequest) error                                                                     // UpdateNote updates an existing note identified by id with the data provided in the DTO.
	GetNotesWithPagination(ctx context.Context, pagination dto.PaginationData) ([]*dto.GetNoteResponse, error)                                        // GetNotesWithPagination retrieves notes with pagination support.
	GetNotesByFilterWithPagination(ctx context.Context, noteFilterData dto.NoteFilter, pagination dto.PaginationData) ([]*dto.GetNoteResponse, error) // GetNotesByFieldWithPagination retrieves notes by a specific field with pagination support.
	GetNoteByID(ctx context.Context, id uint) (dto.GetNoteResponse, error)                                                                            // GetNoteByID retrieves an note by ID.
	DeleteNote(ctx context.Context, id uint) error                                                                                                    // DeleteNoteByID deletes an note by ID.
}

// NewNoteService creates a new instance of noteService.
// This function serves as a constructor for noteService, injecting its dependencies.
func NewNoteService(
	cfg *config.Config,
	noteRepository repository.NoteRepository,
) NoteService {
	return &noteService{
		cfg:            cfg,
		noteRepository: noteRepository,
	}
}

// CreateNote handles the creation of a new note.
// It takes a context for request scoping and a CreateNoteRequest DTO containing the note details.
func (ns *noteService) CreateNote(ctx context.Context, noteDTO dto.CreateNoteRequest) error {
	note := ConvertCreateNoteDTOToModel(noteDTO)
	err := ns.noteRepository.Save(ctx, &note)
	if err != nil && !errors.Is(err, customErrors.ErrDuplicatedKey) {
		ns.cfg.Logger.Error("failed at note creation at CreateNote:", zap.Error(err))
	}
	return err
}

// UpdateNote handles the update of an existing note.
// It requires a note ID and a UpdateNoteRequest DTO containing the updated note details.
func (ns *noteService) UpdateNote(ctx context.Context, id uint, noteDTO dto.UpdateNoteRequest) error {
	note := ConvertUpdateNoteDTOToModel(noteDTO)
	note.ID = id
	err := ns.noteRepository.Save(ctx, &note)
	if err != nil && !errors.Is(err, customErrors.ErrDuplicatedKey) {
		ns.cfg.Logger.Error("failed at note update by ID at UpdateNote:", zap.Uint("id", id), zap.Error(err))
	}
	return err
}

// GetNotesWithPagination retrieves a list of notes with support for pagination.
// It accepts a PaginationData DTO to specify pagination parameters.
func (ns *noteService) GetNotesWithPagination(ctx context.Context, pagination dto.PaginationData) ([]*dto.GetNoteResponse, error) {
	notesResponse, err := ns.GetNotesByFilterWithPagination(ctx, dto.NoteFilter{}, pagination)
	if err != nil {
		if !errors.Is(err, customErrors.ErrRecordNotFound) {
			ns.cfg.Logger.Error("failed at getting notes with pagination in GetNotesWithPagination service:", zap.Any("pagination", pagination), zap.Error(err))
		}
		return notesResponse, err
	}
	return notesResponse, nil
}

// GetNotesByFieldWithPagination retrieves notes based on a specific field, with support for pagination.
// It accepts a FindByFieldData DTO to specify the field and value to search by, along with PaginationData for pagination.
func (ns *noteService) GetNotesByFilterWithPagination(ctx context.Context, noteFilterData dto.NoteFilter, pagination dto.PaginationData) ([]*dto.GetNoteResponse, error) {
	var notesResponse []*dto.GetNoteResponse
	offset := (pagination.Page - 1) * pagination.Limit
	filterModel := ConvertFilterNoteDTOToModel(noteFilterData)
	notes, err := ns.noteRepository.GetByFilterWithPagination(ctx, filterModel, int(offset), int(pagination.Limit))
	if err != nil {
		if !errors.Is(err, customErrors.ErrRecordNotFound) {
			ns.cfg.Logger.Error("failed at getting notes by filter with pagination in GetNotesByFilterWithPagination service:", zap.Any("filter", noteFilterData), zap.Any("pagination", pagination), zap.Error(err))
		}
		return notesResponse, err
	}
	for _, note := range notes {
		noteResponse := ConvertNoteModelToGetNoteDto(*note)
		notesResponse = append(notesResponse, &noteResponse)
	}
	return notesResponse, nil
}

// GetNoteByID retrieves an note by ID.
func (ns *noteService) GetNoteByID(ctx context.Context, id uint) (dto.GetNoteResponse, error) {
	var noteResponse dto.GetNoteResponse
	noteFilterData := dto.NoteFilter{}
	noteFilterData.ID = id
	paginationData := dto.PaginationData{Page: 1, Limit: 10}
	notesResponse, err := ns.GetNotesByFilterWithPagination(ctx, noteFilterData, paginationData)
	if err != nil {
		if !errors.Is(err, customErrors.ErrRecordNotFound) {
			ns.cfg.Logger.Error("failed at getting note by ID in GetNoteByID service:", zap.Uint("id", id), zap.Error(err))
		}
		return noteResponse, err
	}
	if len(notesResponse) == 0 {
		return noteResponse, customErrors.ErrRecordNotFound
	}
	noteResponse = *notesResponse[0]
	return noteResponse, nil
}

// DeleteNoteByID deletes an note by ID.
func (ns *noteService) DeleteNote(ctx context.Context, id uint) error {
	var note model.Note
	note.ID = id
	err := ns.noteRepository.Delete(ctx, &note)
	if err != nil && !errors.Is(err, customErrors.ErrDuplicatedKey) {
		ns.cfg.Logger.Error("failed at deleting note by ID in DeleteNoteByID service:", zap.Uint("id", id), zap.Error(err))
	}
	return err
}

// ConvertCreateNoteDTOToModel converts CreateNoteRequest DTO to Note model.
func ConvertCreateNoteDTOToModel(dto dto.CreateNoteRequest) model.Note {
	return model.Note{
		Title: dto.Title,
		Text:  dto.Text,
	}
}

// ConvertUpdateNoteDTOToModel converts UpdateNoteRequest DTO to Note model.
func ConvertUpdateNoteDTOToModel(dto dto.UpdateNoteRequest) model.Note {
	return model.Note{
		Title: dto.Title,
		Text:  dto.Text,
	}
}

// ConvertNoteModelToGetNoteDto converts Note model to GetNoteResponse DTO.
func ConvertNoteModelToGetNoteDto(model model.Note) dto.GetNoteResponse {
	return dto.GetNoteResponse{
		BaseResponse: dto.BaseResponse{
			ID:        model.ID,
			CreatedAt: model.CreatedAt.UTC().Unix(),
			UpdatedAt: model.UpdatedAt.UTC().Unix(),
		},
		Title: model.Title,
		Text:  model.Text,
	}
}

// ConvertCreateNoteDTOToModel converts CreateNoteRequest DTO to Note model.
func ConvertFilterNoteDTOToModel(dto dto.NoteFilter) model.Note {
	var baseModel model.BaseModel
	baseModel.ID = dto.ID
	if dto.CreatedAt > 0 {
		baseModel.CreatedAt = time.Unix(dto.CreatedAt, 0).UTC()
	}
	if dto.UpdatedAt > 0 {
		baseModel.UpdatedAt = time.Unix(dto.UpdatedAt, 0).UTC()
	}
	return model.Note{
		BaseModel: baseModel,
		Title:     dto.Title,
		Text:      dto.Text,
	}
}
