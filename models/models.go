package models

import (
	"time"
)

/*
gorm.Model definition
type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
*/

// STEP 1: Add the Model.
type UploadedFile struct {
	FileID    string    `gorm:"primaryKey;unique" json:"fileid"`
	Name      string    `json:"name"`
	Size      int64     `json:"size"`
	CreatedAt time.Time `json:"created_at"`
	Type      string
	Isdeleted bool
}
