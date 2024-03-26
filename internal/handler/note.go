package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/ali-shokoohi/notes/internal/config"
	"github.com/ali-shokoohi/notes/internal/constants"
	"github.com/ali-shokoohi/notes/internal/dto"
	customErrors "github.com/ali-shokoohi/notes/internal/errors"
	"github.com/ali-shokoohi/notes/internal/response"
	"github.com/ali-shokoohi/notes/internal/service"
	"github.com/gin-gonic/gin"
)

type noteHandler struct {
	cfg         *config.Config
	noteService service.NoteService
}

type NoteHandler interface {
	CreateNote(c *gin.Context)
	UpdateNote(c *gin.Context)
	GetNotesWithPagination(c *gin.Context)
	GetNotesByFilterWithPagination(c *gin.Context)
	GetNoteByID(c *gin.Context)
	DeleteNote(c *gin.Context)
}

func NewNoteHandler(
	cfg *config.Config,
	noteService service.NoteService,
) NoteHandler {
	return &noteHandler{
		cfg:         cfg,
		noteService: noteService,
	}
}

func (nh *noteHandler) CreateNote(c *gin.Context) {
	var noteData dto.CreateNoteRequest
	err := c.ShouldBind(&noteData)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "binding: "+err.Error())
		return
	}
	err = nh.noteService.CreateNote(c, noteData)
	if err != nil {
		var statusCode int
		if errors.Is(err, customErrors.ErrDuplicatedKey) ||
			errors.Is(err, customErrors.ErrInvalidData) ||
			errors.Is(err, customErrors.ErrInvalidField) ||
			errors.Is(err, customErrors.ErrInvalidValue) ||
			errors.Is(err, customErrors.ErrInvalidValueOfLength) {
			statusCode = http.StatusBadRequest
		} else {
			statusCode = http.StatusInternalServerError
		}
		response.SendError(c, statusCode, err.Error())
		return
	}
	response.SendMessage(c, constants.InsertedSuccessFully)
}

func (nh *noteHandler) UpdateNote(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "invalid id type: "+err.Error())
		return
	}
	var noteData dto.UpdateNoteRequest
	err = c.ShouldBind(&noteData)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "binding: "+err.Error())
		return
	}
	err = nh.noteService.UpdateNote(c, uint(id), noteData)
	if err != nil {
		var statusCode int
		if errors.Is(err, customErrors.ErrDuplicatedKey) ||
			errors.Is(err, customErrors.ErrInvalidData) ||
			errors.Is(err, customErrors.ErrInvalidField) ||
			errors.Is(err, customErrors.ErrInvalidValue) ||
			errors.Is(err, customErrors.ErrInvalidValueOfLength) {
			statusCode = http.StatusBadRequest
		} else {
			statusCode = http.StatusInternalServerError
		}
		response.SendError(c, statusCode, err.Error())
		return
	}
	response.SendMessage(c, constants.UpdatedSuccessFully)
}

func (nh *noteHandler) GetNotesWithPagination(c *gin.Context) {
	var paginationData dto.PaginationData
	err := c.BindQuery(&paginationData)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "binding pagination query: "+err.Error())
		return
	}
	if paginationData.Page == 0 {
		paginationData.Page = 1
	}
	if paginationData.Limit == 0 {
		paginationData.Limit = 5
	}
	notes, err := nh.noteService.GetNotesWithPagination(c, paginationData)
	if err != nil && !errors.Is(err, customErrors.ErrRecordNotFound) {
		var statusCode int
		if errors.Is(err, customErrors.ErrInvalidData) ||
			errors.Is(err, customErrors.ErrInvalidField) ||
			errors.Is(err, customErrors.ErrInvalidValue) ||
			errors.Is(err, customErrors.ErrInvalidValueOfLength) {
			statusCode = http.StatusBadRequest
		} else {
			statusCode = http.StatusInternalServerError
		}
		response.SendError(c, statusCode, err.Error())
		return
	}
	response.SendData(c, notes)
}

func (nh *noteHandler) GetNotesByFilterWithPagination(c *gin.Context) {
	var paginationData dto.PaginationData
	err := c.BindQuery(&paginationData)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "binding pagination query: "+err.Error())
		return
	}
	var noteFilterData dto.NoteFilter
	err = c.BindQuery(&noteFilterData)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "binding find by field query: "+err.Error())
		return
	}
	if paginationData.Page == 0 {
		paginationData.Page = 1
	}
	if paginationData.Limit == 0 {
		paginationData.Limit = 5
	}
	notes, err := nh.noteService.GetNotesByFilterWithPagination(c, noteFilterData, paginationData)
	if err != nil {
		var statusCode int
		if errors.Is(err, customErrors.ErrInvalidData) ||
			errors.Is(err, customErrors.ErrInvalidField) ||
			errors.Is(err, customErrors.ErrInvalidValue) ||
			errors.Is(err, customErrors.ErrRecordNotFound) ||
			errors.Is(err, customErrors.ErrInvalidValueOfLength) {
			statusCode = http.StatusBadRequest
		} else {
			statusCode = http.StatusInternalServerError
		}
		response.SendError(c, statusCode, err.Error())
		return
	}
	response.SendData(c, notes)
}

func (nh *noteHandler) GetNoteByID(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "invalid id type: "+err.Error())
		return
	}
	note, err := nh.noteService.GetNoteByID(c, uint(id))
	if err != nil {
		var statusCode int
		if errors.Is(err, customErrors.ErrInvalidData) ||
			errors.Is(err, customErrors.ErrInvalidField) ||
			errors.Is(err, customErrors.ErrInvalidValue) ||
			errors.Is(err, customErrors.ErrRecordNotFound) ||
			errors.Is(err, customErrors.ErrInvalidValueOfLength) {
			statusCode = http.StatusBadRequest
		} else {
			statusCode = http.StatusInternalServerError
		}
		response.SendError(c, statusCode, err.Error())
		return
	}
	response.SendData(c, note)
}

func (nh *noteHandler) DeleteNote(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "invalid id type: "+err.Error())
		return
	}
	var noteData dto.UpdateNoteRequest
	err = c.ShouldBind(&noteData)
	if err != nil {
		response.SendError(c, http.StatusBadRequest, "binding: "+err.Error())
		return
	}
	err = nh.noteService.DeleteNote(c, uint(id))
	if err != nil {
		var statusCode int
		if errors.Is(err, customErrors.ErrDuplicatedKey) ||
			errors.Is(err, customErrors.ErrInvalidData) ||
			errors.Is(err, customErrors.ErrInvalidField) ||
			errors.Is(err, customErrors.ErrInvalidValue) ||
			errors.Is(err, customErrors.ErrRecordNotFound) ||
			errors.Is(err, customErrors.ErrInvalidValueOfLength) {
			statusCode = http.StatusBadRequest
		} else {
			statusCode = http.StatusInternalServerError
		}
		response.SendError(c, statusCode, err.Error())
		return
	}
	response.SendMessage(c, constants.DeletedSuccessFully)
}
