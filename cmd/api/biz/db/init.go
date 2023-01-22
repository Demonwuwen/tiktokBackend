package db

import (
	"Demonwuwen/tiktokBackend/pkg/consts"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/opentelemetry/logging/logrus"
	"gorm.io/plugin/opentelemetry/tracing"
	"time"
)

var DB *gorm.DB

func Init() {
	var err error
	gormlogrus := logger.New(
		logrus.NewWriter(),
		logger.Config{
			SlowThreshold: time.Millisecond,
			Colorful:      false,
			LogLevel:      logger.Info,
		},
	)
	DB, err = gorm.Open(
		mysql.Open(consts.MySQLDefaultDSN),
		&gorm.Config{
			SkipDefaultTransaction: true,
			PrepareStmt:            false,
			Logger:                 gormlogrus,
		},
	)
	if err != nil {
		panic(err)
	}

	err = DB.Use(tracing.NewPlugin())
	if err != nil {
		panic(err)
	}

}
