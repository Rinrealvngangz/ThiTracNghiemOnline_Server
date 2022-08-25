package repository

import (
	entity "github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/internal/student/entity"
)

func findById(id int) entity.Student {
	student := entity.TestCaseFindByStudent(&entity.Student{})
	return student
}
