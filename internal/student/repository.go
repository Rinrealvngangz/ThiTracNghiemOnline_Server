package student

import (
	"context"

	"github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/internal/student/entity"
	presenter "github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/internal/student/presenter"
)

const CtxUserKey = "userId"

type StudentRepository interface {
	FindById(ctx context.Context, id string) (*entity.Student, error)
	Insert(ctx context.Context, studentRequest presenter.StudentRequest) error
}
