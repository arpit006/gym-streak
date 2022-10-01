package stores

import (
	"fmt"
	"github.com/arpit006/gym_streak/internal/config"
	"github.com/arpit006/gym_streak/internal/logger"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const appDatabaseMigrationsPath = "file://db/migrations"

//createMigrate this function will give you instance of Migrate library to
func createMigrate() *migrate.Migrate {
	dbConf := config.GetDatabaseConfig()
	driver, err := mysql.WithInstance(STORE.db, &mysql.Config{})
	if err != nil {
		logger.Panicf("Error connecting to database to run db migrations on [%s] database: [%s]. Error is - %s", dbConf.Type, dbConf.DatabaseName, err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance(appDatabaseMigrationsPath, dbConf.Type, driver)
	if err != nil {
		logger.Panicf(fmt.Sprintf("Failed to prepare migration on [%s] database: [%s]. Error is - %s", dbConf.Type, dbConf.DatabaseName, err.Error()))
	}

	return m
}

func RunDatabaseMigrations() error {
	m := createMigrate()
	err := m.Up()
	if err != nil {
		if err == migrate.ErrNoChange {
			logger.Info("No changes detected for migration")
			return nil
		}
		logger.Errorf("Migration failed. Error is - %s", err)
	}
	logger.Info("Migration completed successfully")
	return nil
}

func RollbackDatabaseMigrations() error {
	m := createMigrate()
	err := m.Steps(-1)
	if err != nil {
		if err == migrate.ErrNoChange {
			logger.Info("No changes detected for migration")
			return nil
		}
		logger.Errorf("Rollback failed. Error is - %s", err)
	}
	logger.Info("Rollback completed successfully")
	return nil
}
