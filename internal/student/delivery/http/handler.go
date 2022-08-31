package http

import (
	"net/http"

	student "github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/internal/student"
	presenter "github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/internal/student/presenter"
	util "github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/utils"
	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
)

var modulesName string = "MODULE STUDENT"

type studentHandler struct {
	studentUC student.StudentCase
	logger    logging.Logger
}

func NewStudentHandler(studentUC student.StudentCase) *studentHandler {
	return &studentHandler{studentUC: studentUC, logger: *util.Gologger()}
}

func (std *studentHandler) GetStudents(ctx *gin.Context) {
	studentFind := &presenter.StudentFindRequest{}
	ctx.ShouldBindJSON(&studentFind)
	students := std.studentUC.Find(ctx, *studentFind)
	util.APIResponse(ctx, "Find All students success", http.StatusCreated, http.MethodPost, students, modulesName)
}

func (std *studentHandler) GetStudentById(ctx *gin.Context) {
	idStudent := ctx.Param("id")
	student, err := std.studentUC.FindById(ctx, idStudent)
	if err != nil {
		std.logger.Error("Cannot FindById student: "+idStudent, student)
		util.APIResponse(ctx, "FindById studentById  failed", http.StatusInternalServerError, http.MethodGet, nil, modulesName)
	} else {
		util.APIResponse(ctx, "FindById student success", http.StatusOK, http.MethodGet, student, modulesName)
	}

}

func (std *studentHandler) InsertStudent(ctx *gin.Context) {
	student := &presenter.StudentRequest{}
	ctx.ShouldBindJSON(&student)
	resutl, err := util.GoValidator(student)
	if err != nil {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, err.Error(), modulesName)
	}
	if resutl == true {
		result := std.studentUC.Insert(ctx, *student)
		if result != nil {
			std.logger.Error("Cannot insert student:", result)
			util.APIResponse(ctx, "Create new student account failed", http.StatusInternalServerError, http.MethodPost, nil, modulesName)
		} else {
			util.APIResponse(ctx, "Create student success", http.StatusCreated, http.MethodPost, nil, modulesName)
		}

	}
}

func (std *studentHandler) UpdateStudentById(ctx *gin.Context) {
	student := &presenter.StudentUpdateRequest{}
	ctx.ShouldBindJSON(&student)
	idStudent := ctx.Param("id")
	result := std.studentUC.UpdateById(ctx, idStudent, *student)
	if result != nil {
		std.logger.Error("Cannot update student:", result)
		util.APIResponse(ctx, "Update student account failed", http.StatusInternalServerError, http.MethodPut, nil, modulesName)
	} else {
		util.APIResponse(ctx, "Update student success", http.StatusCreated, http.MethodPut, nil, modulesName)
	}
}

func (std *studentHandler) DeleteStudentById(ctx *gin.Context) {
	idStudent := ctx.Param("id")
	result := std.studentUC.DeleteById(ctx, idStudent)
	if result != nil {
		std.logger.Error("Cannot delete student:", result)
		util.APIResponse(ctx, "Delete student account failed", http.StatusInternalServerError, http.MethodDelete, nil, modulesName)
	} else {
		util.APIResponse(ctx, "Delete student success", http.StatusCreated, http.MethodDelete, nil, modulesName)
	}
}
