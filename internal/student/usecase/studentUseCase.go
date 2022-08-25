package usecase

import (
	student "github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/internal/student"
)

type studentUseCase struct {
	studentRepo student.StudentRepository
}

func studentFindByIdUserCase(studentRepo student.StudentRepository) *studentUseCase {
	return &studentUseCase{
		studentRepo: studentRepo,
	}
}
