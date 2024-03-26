package model

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	Title string
	Text  string
}
