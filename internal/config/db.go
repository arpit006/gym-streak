package config

import "fmt"

type DatabaseConfig struct {
	Host               string
	Port               int
	Username           string
	Password           string
	MaxOpenConnections int
	MaxIdleConnections int
	ConnMaxLifetime    int
	ConnMaxIdleTime    int
	DatabaseName       string
	Protocol           string
	DSN                string
	Type               string
}

func readDatabaseConfig() DatabaseConfig {
	dbConf := DatabaseConfig{
		Host:               readViperString("db_host"),
		Port:               readViperInteger("db_port"),
		Username:           readViperString("db_user"),
		Password:           readViperString("db_pass"),
		MaxOpenConnections: readViperInteger("db_max_open_connections"),
		MaxIdleConnections: readViperInteger("db_max_idle_connections"),
		ConnMaxLifetime:    readViperInteger("db_conn_max_lifetime"),
		ConnMaxIdleTime:    readViperInteger("db_conn_max_idle_time"),
		DatabaseName:       readViperString("db_name"),
		Protocol:           readViperString("db_protocol"),
		Type:               readViperString("db_type"),
	}
	dbConf.DSN = fmt.Sprintf("%s:%s@%s(%s:%d)/%s", dbConf.Username, dbConf.Password, dbConf.Protocol, dbConf.Host, dbConf.Port, dbConf.DatabaseName)
	return dbConf
}
