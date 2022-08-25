package server

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	g *gin.Engine
}

func NewServer() *Server {
	return &Server{g: gin.New()}
}

func (s *Server) Run() error {
	s.MapHandlers(*s.g)
	s.g.Run(":8080")
	return nil
}
