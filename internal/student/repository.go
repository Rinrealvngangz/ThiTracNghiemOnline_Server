package student

import (
	"context"

	"github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/internal/student/entity"
	presenter "github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/internal/student/presenter"
)

const CtxUserKey = "userId"

type StudentRepository interface {
	Find(ctx context.Context, studentFindRequest presenter.StudentFindRequest) *presenter.StudentFindResponse
	FindById(ctx context.Context, id string) (*entity.Student, error)
	Insert(ctx context.Context, studentRequest presenter.StudentRequest) error
	UpdateById(ctx context.Context, id string, studentRequest presenter.StudentUpdateRequest) error
	DeleteById(ctx context.Context, id string) error
	LoginByPhone(ctx context.Context, phoneNumber string, password string) (*presenter.StudentResponse, error)
}
