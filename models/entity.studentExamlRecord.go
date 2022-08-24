package model

import (
	"time"
)

type EntityStudentExamRecord struct {
	StudentExamRecordId string `gorm:"primaryKey;"`
	StudentId           string `gorm:"type:varchar(255);unique;not null"`  // id thi sinh
	ExamId              string `gorm:"type:varchar(255);unique;not null"`  // id de thi
	Note                string `gorm:"type:varchar(1000);unique;not null"` // ghi chu
	Active              bool   `gorm:"type:bool;default:true"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	IsDeleted           bool `gorm:"type:bool;default:true"`
}
