package server

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	g  *gin.Engine
	db *gorm.DB
}

func NewServer(db *gorm.DB) *Server {
	return &Server{g: gin.New(), db: db}
}

func (s *Server) Run() error {
	s.MapHandlers(*s.g)
	s.g.Run(":8080")
	return nil
}
