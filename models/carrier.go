package models

import (
	"gorm.io/gorm"
)


type Carrier struct {
	gorm.Model
	ID uint `gorm:"primaryKey;autoIncrement"`
    Name string
}