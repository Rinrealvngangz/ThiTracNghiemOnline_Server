package usecase

import (
	"github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/internal/student"
	entity "github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/internal/student/entity"
)

type studentUseCase struct {
	studentRepo student.StudentRepository
}

func NewStudentUseCase(studentRepo student.StudentRepository) student.StudentCase {
	return studentUseCase{
		studentRepo: studentRepo,
	}
}

// FindById implements student.StudentCase
func (std studentUseCase) FindById(id int) *entity.Student {
	return std.studentRepo.FindById(1)
}
