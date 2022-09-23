package db

import (
	"fmt"

	"github.com/LambdaTest/photon/config"
	"github.com/LambdaTest/photon/pkg/global"
	"github.com/LambdaTest/photon/pkg/lumber"
	"github.com/LambdaTest/photon/pkg/models"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// TODO: use read replica

// Connect create connection with database
func Connect(cfg *config.Config, logger lumber.Logger) (models.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", cfg.DB.User, cfg.DB.Password, "tcp", cfg.DB.Host, cfg.DB.Port, cfg.DB.Name)
	db, err := sqlx.Connect("mysql", connectionString+"?parseTime=true&charset=utf8mb4")
	if err != nil {
		return nil, err
	}
	logger.Infof("Database connected successfully")
	db.SetMaxIdleConns(global.MysqlMaxIdleConnection)
	db.SetMaxOpenConns(global.MysqlMaxOpenConnection)
	db.SetConnMaxLifetime(global.MysqlMaxConnectionLifetime)

	return &DB{conn: db, logger: logger}, nil
}
