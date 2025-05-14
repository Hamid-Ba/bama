package models

import "time"

type BaseModel struct {
	Id         int       `gorm:"primarykey"`
	Created_at time.Time `gorm:"type:TIMESTAMP with time zone;not null"`
	Updated_at time.Time `gorm:"type:TIMESTAMP with time zone;null"`
	IsActive   bool
}

func (base_model *BaseModel) BeforeCreate() {
	base_model.Created_at = time.Now().UTC()
}
