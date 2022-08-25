package entity

import (
	"time"

	"github.com/google/uuid"
)

type Student struct {
	idStudent   string `gorm:"primaryKey;"`
	fullName    string `gorm:"type:varchar(255);unique;not null"` // tên thi sinh
	phoneNumber string `gorm:"type:varchar(255);null"`            // sdt
	email       string `gorm:"type:varchar(255);null"`            // loại đề
	password    string `gorm:"type:varchar(255);not null"`
	createdAt   time.Time
	updatedAt   time.Time
	isDeleted   int
}

func TestCaseFindByStudent() *Student {
	var student Student
	student.idStudent = uuid.New().String()
	student.fullName = "Nguyen Mau Tuan"
	student.createdAt = time.Now().Local()
	student.updatedAt = time.Now().Local()
	student.email = "tuan@gmail.com"
	student.phoneNumber = "0949239777"
	student.password = "string"
	return &student
}
