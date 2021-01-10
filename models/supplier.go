package models

import (
	"gorm.io/gorm"
)


type Supplier struct {
	gorm.Model
	ID uint `gorm:"primaryKey;autoIncrement"`
    Name string
}