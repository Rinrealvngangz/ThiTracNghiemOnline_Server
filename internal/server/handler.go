package server

import (
	http "github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/internal/student/delivery/http"
	repository "github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/internal/student/repository"
	usecase "github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/internal/student/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (s *Server) MapHandlers(r gin.Engine, db *gorm.DB) error {

	// repos
	studentRepository := repository.NewStudenRepository(s.db)

	studentUseCase := usecase.NewStudentUseCase(studentRepository)
	//handler
	studentHandler := http.NewStudentHandler(studentUseCase)

	//versioning
	v1 := r.Group("/api/v1")

	studentGroup := v1.Group("/students")

	http.MapAuthRoutes(*studentGroup, studentHandler, db)
	return nil
}
