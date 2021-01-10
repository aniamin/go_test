package models

import (
	"gorm.io/gorm"
)


type CarrierBeanType struct {
	gorm.Model
	ID uint `gorm:"primaryKey;autoIncrement"`
	Carrier_id int
	Bean_type_id  int
}