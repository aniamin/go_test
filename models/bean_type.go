package models

import (
	"gorm.io/gorm"
)

type BeanType struct {
	gorm.Model
	ID uint `gorm:"primaryKey;autoIncrement"`
    Name string
}