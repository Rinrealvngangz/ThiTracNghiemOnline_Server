package entity

import (
	"time"

	"github.com/google/uuid"
)

type Student struct {
	IdStudent   string `gorm:"primaryKey;auto_increment;not_null"`
	FullName    string `gorm:"type:varchar(255);not null"`
	PhoneNumber string `gorm:"type:varchar(255);unique;null"`
	Email       string `gorm:"type:varchar(255);unique;null"`
	Password    string `gorm:"type:varchar(255);not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	IsDeleted   int
}

func TestCaseFindByStudent() *Student {
	var student Student
	student.IdStudent = uuid.New().String()
	student.FullName = "Nguyen Mau Tuan"
	student.CreatedAt = time.Now().Local()
	student.UpdatedAt = time.Now().Local()
	student.Email = "tuan@gmail.com"
	student.PhoneNumber = "0949239777"
	student.Password = "string"
	return &student
}

func (student *Student) BeforeCreate() error {
	student.IdStudent = uuid.New().String()
	student.CreatedAt = time.Now().Local()
	return nil
}
