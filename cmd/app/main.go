package main

import (
	config "github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/configs"
	"github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/db"
	server "github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/internal/server"
	util "github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/utils"
)

func main() {
	logger := util.Gologger()
	cfg := config.NewConfig()
	db := db.GetPostgresInstance(cfg, false)

	logger.Info("Starting api server")
	s := server.NewServer(db, cfg.Port)
	if err := s.Run(); err != nil {
		logger.Fatal(err)
	}
}
