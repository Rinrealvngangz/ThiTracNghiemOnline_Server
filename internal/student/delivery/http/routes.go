package http

import (
	student "github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/internal/student"
	util "github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Map auth routes
func MapAuthRoutes(router gin.RouterGroup, h student.Handler, db *gorm.DB) {
	router.POST("/register", h.InsertStudent)
	router.POST("/loginByPhone", h.LoginByPhone)

	router.Use(util.JwtAuthMiddleware())
	router.Use(util.JwtAuthUser(db))
	router.GET("/findById/:id", h.GetStudentById)
	router.POST("/find", h.GetStudents)
	router.PUT("/updateById/:id", h.UpdateStudentById)
	router.POST("/insert", h.InsertStudent)
	router.DELETE("/deleteById/:id", h.DeleteStudentById)

}
