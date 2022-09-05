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

func (std studentUseCase) FindById(ctx context.Context, id string) (*entity.Student, error) {
	return std.studentRepo.FindById(ctx, id)
}

func (std studentUseCase) Insert(ctx context.Context, studentRequest presenter.StudentRequest) error {
	return std.studentRepo.Insert(ctx, studentRequest)
}

func (std studentUseCase) UpdateById(ctx context.Context, id string, studentRequest presenter.StudentUpdateRequest) error {
	return std.studentRepo.UpdateById(ctx, id, studentRequest)
}

func (std studentUseCase) DeleteById(ctx context.Context, id string) error {
	return std.studentRepo.DeleteById(ctx, id)
}

func (std studentUseCase) Find(ctx context.Context, studentFindRequest presenter.StudentFindRequest) *presenter.StudentFindResponse {
	return std.studentRepo.Find(ctx, studentFindRequest)
}

func (std studentUseCase) LoginByPhone(ctx context.Context, phoneNumber string, password string) (*presenter.StudentResponse, error) {
	return std.studentRepo.LoginByPhone(ctx, phoneNumber, password)
}
