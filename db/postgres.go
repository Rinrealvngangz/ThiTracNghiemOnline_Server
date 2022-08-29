package db

import (
	config "github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/configs"
	entity "github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/internal/student/entity"
	util "github.com/Rinrealvngangz/ThiTracNghiemOnline_Server/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetPostgresInstance(cfg *config.Configuration, migrate bool) *gorm.DB {
	logger := util.Gologger()
	dsn := cfg.DatabaseConnectionURL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.Error(err)
		panic(err)
	}

	if migrate {
		db.AutoMigrate(&entity.Student{})
	}
	return db
}
