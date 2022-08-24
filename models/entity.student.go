package model

import (
	"time"
)

type EntityStudents struct {
	StudentId   string `gorm:"primaryKey;"`
	FullName    string `gorm:"type:varchar(255);unique;not null"` // tên thi sinh
	BirthDate   string `gorm:"type:varchar(255);unique;not null"` // ngay sinh
	PhoneNumber string `gorm:"type:varchar(255);not null"`        // sdt
	Email       string `gorm:"type:varchar(255);not null"`        // loại đề
	Active      bool   `gorm:"type:bool;default:true"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	IsDeleted   bool `gorm:"type:bool;default:true"`
}
