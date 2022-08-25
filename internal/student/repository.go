package student

import (
	entity "github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/internal/student/entity"
)

const CtxUserKey = "userId"

type StudentRepository interface {
	FindById(id int) *entity.Student
}
