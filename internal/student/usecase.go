package student

import (
	entity "github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/internal/student/entity"
)

type StudentCase interface {
	FindById(id int) *entity.Student
}
