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
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: GetConnectionString(),
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetConnectionString() string {
	env := config.GetCofig()

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable&timezone=Asia/Shanghai",
		env.POSTGRES_USER,
		env.POSTGRES_PASSWORD,
		env.POSTGRES_HOST,
		env.POSTGRES_PORT,
		env.POSTGRES_DB,
	)
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
