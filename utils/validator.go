package util

import (
	"github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/internal/student/presenter"
	"github.com/asaskevich/govalidator"
)

// Read request body and validate
func GoValidator(studentRequest *presenter.StudentRequest) (bool, error) {
	result, err := govalidator.ValidateStruct(studentRequest)
	return result, err
}
