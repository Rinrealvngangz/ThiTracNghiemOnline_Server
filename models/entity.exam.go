package model

import (
	"time"
)

type EntityExams struct {
	ExamId    string `gorm:"primaryKey;"`
	ExamName  string `gorm:"type:varchar(255);unique;not null"` // tên môn thi
	Time      string `gorm:"type:varchar(255);unique;not null"` // thời gian
	Subject   string `gorm:"type:varchar(255);not null"`        // môn thi
	ExameType string `gorm:"type:varchar(255);not null"`        // loại đề
	Active    bool   `gorm:"type:bool;default:true"`
	Grade     string `gorm:"type:varchar(255);not null"` // lớp
	CreatedAt time.Time
	UpdatedAt time.Time
	IsDeleted bool `gorm:"type:bool;default:true"`
}
