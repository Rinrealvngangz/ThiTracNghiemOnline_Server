package repository

import (
	"context"

	student "github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/internal/student"
	entity "github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/internal/student/entity"
	presenter "github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/internal/student/presenter"
	"gorm.io/gorm"
)

type studentRepository struct {
	db *gorm.DB
}

func NewStudenRepository(db *gorm.DB) student.StudentRepository {
	return &studentRepository{db: db}
}

func (std *studentRepository) Find(ctx context.Context, studentFindRequest presenter.StudentFindRequest) *presenter.StudentFindResponse {
	students := []entity.Student{}
	queryBuilder := std.db.WithContext(ctx)
	if studentFindRequest.SearchText != "" {
		searchText := studentFindRequest.SearchText
		queryBuilder = queryBuilder.Where("email = ?", searchText).
			Or("phone_number = ?", searchText).
			Or("full_name = ?", searchText).Find(&students)

	}
	if studentFindRequest.Limit == 0 {
		studentFindRequest.Limit = 10
	}
	if studentFindRequest.Offset != 0 {
		queryBuilder = queryBuilder.Offset(studentFindRequest.Offset)
	}
	queryBuilder = queryBuilder.Where("is_deleted != 1")
	queryBuilder.Limit(studentFindRequest.Limit).Find(&students)
	count := len(students)
	responseStudents := presenter.StudentFindResponse{
		Students: &students,
		Count:    count,
	}
	return &responseStudents
}

func (std *studentRepository) FindById(ctx context.Context, id string) (*entity.Student, error) {
	student := entity.Student{IdStudent: id}
	result := std.db.WithContext(ctx).Where("is_deleted != 1").First(&student)
	return &student, result.Error
}

func (std *studentRepository) Insert(ctx context.Context, studentRequest presenter.StudentRequest) error {
	student := new(entity.Student)
	student.FullName = studentRequest.FullName
	student.Email = studentRequest.Email
	student.PhoneNumber = studentRequest.PhoneNumber
	student.Password = studentRequest.Password
	student.BeforeCreate()
	result := std.db.WithContext(ctx).Create(&student)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (std *studentRepository) UpdateById(ctx context.Context, id string, studentRequest presenter.StudentUpdateRequest) error {
	student := entity.Student{IdStudent: id}
	studentExist := std.db.WithContext(ctx).First(&student)
	if studentExist.Error != nil {
		return studentExist.Error
	} else {
		std.db.Model(&student).Updates(studentRequest)
		return nil
	}
}

// DeleteById implements student.StudentRepository
func (std *studentRepository) DeleteById(ctx context.Context, id string) error {
	student := entity.Student{IdStudent: id}
	studentExist := std.db.WithContext(ctx).First(&student)
	if studentExist.Error != nil {
		return studentExist.Error
	} else {
		std.db.Model(&student).Where("id_student = ?", id).Update("is_deleted", 1)
		return nil
	}
}
