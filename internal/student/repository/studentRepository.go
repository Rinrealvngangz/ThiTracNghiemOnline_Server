package repository

import (
	"context"

	student "github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/internal/student"
	entity "github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/internal/student/entity"
	presenter "github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/internal/student/presenter"
	"gorm.io/gorm"
)

type studentRepository struct {
	db *gorm.DB
}

func NewStudenRepository(db *gorm.DB) student.StudentRepository {
	return &studentRepository{db: db}
}

// FindById implements student.StudentRepository
func (studentRepository) FindById(id int) *entity.Student {
	student := entity.TestCaseFindByStudent()
	return student
}

func (std *studentRepository) Insert(ctx context.Context, studentRequest presenter.StudentRequest) error {
	student := new(entity.Student)
	student.FullName = studentRequest.FullName
	student.Email = studentRequest.Email
	student.PhoneNumber = studentRequest.PhoneNumber
	student.Password = studentRequest.Password
	student.BeforeCreate()
	result := std.db.WithContext(ctx).Create(&student)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
