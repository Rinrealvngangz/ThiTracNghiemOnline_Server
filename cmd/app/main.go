package main

import (
	server "github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/internal/server"
	util "github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/utils"
)

func main() {
	logger := util.Gologger()
	logger.Info("Starting api server")
	s := server.NewServer()
	if err := s.Run(); err != nil {
		logger.Fatal(err)
	}
}
