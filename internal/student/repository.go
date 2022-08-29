package student

import (
	"context"

	entity "github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/internal/student/entity"
	presenter "github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/internal/student/presenter"
)

const CtxUserKey = "userId"

type StudentRepository interface {
	FindById(id int) *entity.Student
	Insert(ctx context.Context, studentRequest presenter.StudentRequest) error
}
