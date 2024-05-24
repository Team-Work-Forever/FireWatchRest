package adapters

import (
	"fmt"

	"github.com/Team-Work-Forever/FireWatchRest/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var database *gorm.DB

func newDatabaseGorm() (*gorm.DB, error) {
	env := config.GetCofig()
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		env.POSTGRES_HOST,
		env.POSTGRES_USER,
		env.POSTGRES_PASSWORD,
		env.POSTGRES_DB,
		env.POSTGRES_PORT,
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetDatabase() *gorm.DB {
	var err error

	if database == nil {
		database, err = newDatabaseGorm()

		if err != nil {
			panic(err)
		}
	}

	return database
}
