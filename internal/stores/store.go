package stores

import (
	"database/sql"
	"fmt"
	"github.com/arpit006/gym_streak/internal/config"
	"github.com/arpit006/gym_streak/internal/logger"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Store struct {
	db *sql.DB
}

var STORE *Store = nil

func Init() {
	dbConf := config.GetDatabaseConfig()
	//dataSourceName := fmt.Sprintf("%s:%s@/%s", dbConf.Username, dbConf.Password, dbConf.DatabaseName)
	db, err := sql.Open(dbConf.Type, dbConf.DSN)
	if err != nil {
		logger.Error(err.Error())
		logger.Panic(fmt.Sprintf("Could not connect to [%s] database [%s]", dbConf.Type, dbConf.DSN))
	}

	db.SetConnMaxLifetime(time.Duration(dbConf.ConnMaxLifetime))
	db.SetConnMaxIdleTime(time.Duration(dbConf.ConnMaxIdleTime))
	db.SetMaxOpenConns(dbConf.MaxOpenConnections)
	db.SetMaxIdleConns(dbConf.MaxIdleConnections)

	e := db.Ping()
	if e != nil {
		logger.Error(e.Error())
		logger.Panic("STORE connection was un-successful......")
	} else {
		logger.Info("######## STORE connection successful....")
	}

	STORE = &Store{
		db: db,
	}
}
