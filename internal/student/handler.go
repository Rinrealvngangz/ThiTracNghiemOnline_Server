package student

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetStudentById(ctx *gin.Context)
	InsertStudent(ctx *gin.Context)
	UpdateStudentById(ctx *gin.Context)
	DeleteStudentById(ctx *gin.Context)
}
