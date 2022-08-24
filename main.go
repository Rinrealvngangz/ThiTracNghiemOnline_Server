package main

import (
	util "github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	logger := util.Gologger()
	router := SetupRouter()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"title":   "GoLang",
		})
	})
	logger.Info("Server is running...")
	logger.Info("OK")
	router.Run(":" + util.GodotEnv("GO_PORT"))
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))
	return router
}
