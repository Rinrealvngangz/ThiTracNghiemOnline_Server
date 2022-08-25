package entity

import (
	"time"

	"github.com/google/uuid"
)

type Student struct {
	IdStudent   string `gorm:"primaryKey;"`
	FullName    string `gorm:"type:varchar(255);unique;not null"` // tên thi sinh
	PhoneNumber string `gorm:"type:varchar(255);null"`            // sdt
	Email       string `gorm:"type:varchar(255);null"`            // loại đề
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
