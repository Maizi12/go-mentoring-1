package db

import (
	"fmt"
	"go-mentoring-1/app"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	grmsql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connection() *gorm.DB {
	configuration := app.GetConfig()

	newLogger :=
		logger.New(

			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				LogLevel:                  logger.Info,
				IgnoreRecordNotFoundError: true,  // Ignore ErrRecordNotFound error for logger
				Colorful:                  false, // Enable colorful output// Log level: logger.Silent, logger.Error, logger.Warn, logger.Info
			},
		)

	dbUser := configuration.Database.Username
	dbPassword := configuration.Database.Password
	dbURL := configuration.Database.Host
	dbName := configuration.Database.Dbname
	SetIdle := configuration.Database.MaxIdleConns
	SetOpen := configuration.Database.MaxOpenConns
	SetLifetime := configuration.Database.MaxLifetime

	addr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbURL, "3306", dbName)
	db, err := gorm.Open(grmsql.Open(addr), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Println("err connect db main:", err)
	}

	sqlDB, errSet := db.DB()
	if errSet != nil {
		log.Fatal("Failed to get database object from GORM:", err)
	}

	sqlDB.SetMaxIdleConns(SetIdle)
	sqlDB.SetMaxOpenConns(SetOpen)
	sqlDB.SetConnMaxLifetime(time.Duration(SetLifetime))

	return db

}
