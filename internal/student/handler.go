package student

import "github.com/gin-gonic/gin"

type Handler interface {
	GetStudent() gin.HandlerFunc
}
