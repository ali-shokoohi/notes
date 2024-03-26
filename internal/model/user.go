package model

type Note struct {
	BaseModel
	Title string
	Text  string
}

func (u Note) Table() string {
	return "notes"
}
