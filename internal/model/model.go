package model

import (
	"time"

	"gorm.io/gorm"
)

type DBModel interface {
	Table() string
}

type BaseModel struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	// TODO: Its better to have our own deletedAt (For multi ORM)
	DeletedAt gorm.DeletedAt `gorm:"index"` //add soft delete in gorm

}
