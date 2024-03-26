package dto

type CreateNoteRequest struct {
	Title string `json:"title"`
	Text  string `json:"text" binding:"required"`
}

type UpdateNoteRequest struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type GetNoteResponse struct {
	BaseResponse
	Title string `json:"title" form:"title"`
	Text  string `json:"text" form:"text"`
}

type NoteFilter struct {
	GetNoteResponse
}
