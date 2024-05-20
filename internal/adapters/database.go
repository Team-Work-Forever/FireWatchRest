package adapters

import (
	"fmt"

	"github.com/Team-Work-Forever/FireWatchRest/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database interface {
	GetDatabase() *gorm.DB
}

type DatabaseGorm struct {
	database *gorm.DB
}

func NewDatabaseGorm() (*DatabaseGorm, error) {
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
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		return nil, err
	}

	return &DatabaseGorm{
		database: db,
	}, nil
}

func (db *DatabaseGorm) GetDatabase() *gorm.DB {
	return db.database
}
