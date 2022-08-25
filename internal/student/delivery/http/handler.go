package http

import (
	"log"
	"net/http"

	student "github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/internal/student"
	"github.com/gin-gonic/gin"
)

type studentHandler struct {
	studentUC student.StudentCase
}

type Test struct {
	test string
}

type Responses struct {
	StatusCode int         `json:"statusCode"`
	Method     string      `json:"method"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}

func NewStudentHandler(studentUC student.StudentCase) *studentHandler {
	return &studentHandler{studentUC: studentUC}
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
		log.Println(student)
		c.JSON(http.StatusOK, jsonResponse)

	}
}
