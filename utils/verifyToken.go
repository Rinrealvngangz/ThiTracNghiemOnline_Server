package util

import (
	"github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/internal/student/entity"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func verifyTokenStudent(db *gorm.DB, c *gin.Context) (string, error) {
	idStudent, err := ExtractTokenID(c)
	if err != nil {
		return "", err
	} else {

		student := entity.Student{IdStudent: idStudent}
		studentExist := db.WithContext(c).First(&student)
		if studentExist.Error != nil {
			return "", studentExist.Error
		} else {
			return idStudent, nil
		}
	}

}
