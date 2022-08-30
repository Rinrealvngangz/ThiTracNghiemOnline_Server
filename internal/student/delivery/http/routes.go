package http

import (
	student "github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/internal/student"
	"github.com/gin-gonic/gin"
)

// Map auth routes
func MapAuthRoutes(router gin.RouterGroup, h student.Handler) {
	router.GET("/:id", h.GetStudentById)
	router.POST("/insert", h.InsertStudent)
}
