package util

import (
	"github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/internal/student/presenter"
	"github.com/asaskevich/govalidator"
)

func GoValidatorInsert(studentRequest *presenter.StudentRequest) (bool, error) {
	result, err := govalidator.ValidateStruct(studentRequest)
	return result, err
}

func GoValidatorLogin(studentRequest *presenter.StudentLoginRequest) (bool, error) {
	result, err := govalidator.ValidateStruct(studentRequest)
	return result, err
}
