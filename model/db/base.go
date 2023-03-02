package db

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID         int       `gorm:"primary_key" json:"id"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

func (model *BaseModel) BeforeCreate(db *gorm.DB) error {
	model.CreateTime = time.Now()
	model.UpdateTime = time.Now()
	return nil
}

func (model *BaseModel) BeforeUpdate(db *gorm.DB) error {
	model.UpdateTime = time.Now()
	return nil
}
