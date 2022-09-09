package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}

func JwtAuthUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStudent, err := verifyTokenStudent(db, c)
		Gologger().Info(idStudent)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized user")
			c.Abort()
			return
		}
		c.Next()
	}
}
