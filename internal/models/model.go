package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Note struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid"`
	Title    string
	SubTitle string
	Text     string
}

type User struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid"`
	// ID       uuid.UUID `gor:"uuid.Nype:uuid"`
	Name     string
	Email    string `gorm:"unique"`
	Password []byte
}
