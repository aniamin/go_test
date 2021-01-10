package models

import (
	"gorm.io/gorm"
)


type SupplierBeanType struct {
	gorm.Model
	ID uint `gorm:"primaryKey;autoIncrement"`
	SupplierId int
	BeanTypeId  int
}