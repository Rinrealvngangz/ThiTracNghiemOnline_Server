package http

import (
	student "github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/internal/student"
	"github.com/gin-gonic/gin"
)

// Map auth routes
func MapAuthRoutes(router gin.RouterGroup, h student.Handler) {
	router.POST("/find", h.GetStudents)
	router.GET("/findById/:id", h.GetStudentById)
	router.POST("/insert", h.InsertStudent)
	router.PUT("/updateById/:id", h.UpdateStudentById)
	router.DELETE("/deleteById/:id", h.DeleteStudentById)
}
