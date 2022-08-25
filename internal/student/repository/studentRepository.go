package repository

import (
	student "github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/internal/student"
	entity "github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/internal/student/entity"
)

type studentRepository struct {
}

func NewStudenRepository() student.StudentRepository {
	return studentRepository{}
}

// FindById implements student.StudentRepository
func (studentRepository) FindById(id int) *entity.Student {
	student := entity.TestCaseFindByStudent()
	return student
}
