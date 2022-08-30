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

// GetUser implements student.Handler
func (std *studentHandler) GetStudentById(ctx *gin.Context) {
	idStudent := ctx.Param("id")
	//idStudent := uuid.Must(uuid.FromString(paramId))

	student, err := std.studentUC.FindById(ctx, idStudent)
	std.logger.Info(student)
	if err != nil {
		std.logger.Error("Cannot find student: "+idStudent, student)
		util.APIResponse(ctx, "Find studentById  failed", http.StatusInternalServerError, http.MethodGet, nil, modulesName)
	} else {
		util.APIResponse(ctx, "Find student success", http.StatusOK, http.MethodGet, student, modulesName)
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
