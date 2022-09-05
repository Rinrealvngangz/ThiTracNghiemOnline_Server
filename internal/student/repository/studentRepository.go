package repository

import (
	"context"
	"encoding/json"

	"errors"

	student "github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/internal/student"
	entity "github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/internal/student/entity"
	presenter "github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/internal/student/presenter"
	util "github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/utils"
	"gorm.io/gorm"
)

type studentRepository struct {
	db *gorm.DB
}

func NewStudenRepository(db *gorm.DB) student.StudentRepository {
	return &studentRepository{db: db}
}

func (std *studentRepository) Find(ctx context.Context, studentFindRequest presenter.StudentFindRequest) *presenter.StudentFindResponse {

	student := []entity.Student{}
	studentResponse := []presenter.StudentResponse{}
	queryBuilder := std.db.WithContext(ctx)
	if studentFindRequest.SearchText != "" {
		searchText := studentFindRequest.SearchText
		queryBuilder = queryBuilder.Where("email = ?", searchText).
			Or("phone_number = ?", searchText).
			Or("full_name = ?", searchText).Find(&student)
	}
	if studentFindRequest.Limit == 0 {
		studentFindRequest.Limit = 10
	}
	if studentFindRequest.Offset != 0 {
		queryBuilder = queryBuilder.Offset(studentFindRequest.Offset)
	}
	queryBuilder = queryBuilder.Where("is_deleted != 1")
	queryBuilder.Limit(studentFindRequest.Limit).Find(&student)
	count := len(student)

	studentRecord, _ := json.Marshal(student)
	json.Unmarshal([]byte(studentRecord), &studentResponse)

	responseStudents := presenter.StudentFindResponse{
		Students: &studentResponse,
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
	student.Password, _ = util.HashPassword(studentRequest.Password)
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

func (std *studentRepository) LoginByPhone(ctx context.Context, phoneNumber string, password string) (*presenter.StudentResponse, error) {
	queryBuilder := std.db.WithContext(ctx).Where("is_deleted != 1").Where("phone_number = ?", phoneNumber)
	student := entity.Student{}
	studentResponse := presenter.StudentResponse{}
	isExistStudent := queryBuilder.First(&student)

	if isExistStudent.Error != nil {
		return &studentResponse, errors.New("PhoneNumber is not exist")
	} else {
		match := util.CheckPasswordHash(password, student.Password)
		if match == true {
			result := queryBuilder.Where("password = ?", student.Password).First(&studentResponse)
			if result.Error != nil {
				return &studentResponse, result.Error
			}
			return &studentResponse, nil
		}
		return &studentResponse, errors.New("Password is not correct")
	}
}
