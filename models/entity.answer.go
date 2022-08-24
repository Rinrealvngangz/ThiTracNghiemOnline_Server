package model

import (
	"time"
)

type EntityAnswers struct {
	AnswerId    string `gorm:"primaryKey;"`
	AnswersName string `gorm:"type:varchar(255);unique;not null"` // cau tra loi
	QuestionId  string `gorm:"type:varchar(255);unique;not null"` // id cau hoi
	IsCorrect   bool   `gorm:"type:bool;default:false"`
	Active      bool   `gorm:"type:bool;default:true"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	IsDeleted   bool `gorm:"type:bool;default:true"`
}
