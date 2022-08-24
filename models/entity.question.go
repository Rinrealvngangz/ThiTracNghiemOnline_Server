package model

import (
	"time"
)

type EntityQuestions struct {
	QuestionId   string `gorm:"primaryKey;"`
	QuestionName string `gorm:"type:varchar(255);unique;not null"` // ten cau hoi
	Active       bool   `gorm:"type:bool;default:true"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	IsDeleted    bool `gorm:"type:bool;default:true"`
}
