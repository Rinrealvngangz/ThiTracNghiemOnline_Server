package server

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	g    *gin.Engine
	db   *gorm.DB
	port string
}

func NewServer(db *gorm.DB, port string) *Server {
	return &Server{g: gin.New(), db: db, port: port}
}

func (s *Server) Run() error {
	s.MapHandlers(*s.g, *&s.db)
	s.g.Run(":" + s.port)
	return nil
}
