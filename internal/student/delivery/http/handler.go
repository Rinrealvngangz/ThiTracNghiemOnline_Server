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

type Responses struct {
	StatusCode int         `json:"statusCode"`
	Method     string      `json:"method"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func NewStudentHandler(studentUC student.StudentCase) *studentHandler {
	return &studentHandler{studentUC: studentUC, logger: *util.Gologger()}
}

// GetUser implements student.Handler
func (std *studentHandler) GetStudent() gin.HandlerFunc {

	return func(c *gin.Context) {
		var idStudent int = 1
		student := std.studentUC.FindById(idStudent)
		jsonResponse := Responses{
			StatusCode: http.StatusOK,
			Method:     http.MethodGet,
			Message:    "Success",
			Data:       student,
		}
		c.JSON(http.StatusOK, jsonResponse)
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
