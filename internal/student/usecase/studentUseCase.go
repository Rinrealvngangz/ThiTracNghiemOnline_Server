package usecase

import (
	"context"

	"github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/internal/student"
	"github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/internal/student/entity"
	presenter "github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/internal/student/presenter"
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
func (std studentUseCase) FindById(ctx context.Context, id string) (*entity.Student, error) {
	return std.studentRepo.FindById(ctx, id)
}

// Insert implements student.StudentCase
func (std studentUseCase) Insert(ctx context.Context, studentRequest presenter.StudentRequest) error {
	return std.studentRepo.Insert(ctx, studentRequest)
}
