package util

import (
	"context"

	entity "github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/internal/student/entity"
	"gorm.io/gorm"
)

func verifyTokenStudent(db *gorm.DB, ctx context.Context, idStudent string) {
	student := entity.Student{IdStudent: idStudent}
	studentExist := db.WithContext(ctx).First(&student)

}
