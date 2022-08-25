package mysql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/arymaulanamalik/bobobox_test/pkg/logger"
	stg "github.com/arymaulanamalik/bobobox_test/repository/storages"
)

type MysqlConfig struct {
	URL      string
	Schema   string
	User     string
	Password string
}

type dbMysql struct {
	db *sql.DB
}

func NewMysql(cfg MysqlConfig) (stg.ReadWriter, error) {

	logger.Info(" - Init MySQL Database Connection.", nil)
	log.Println(cfg)

	schemaURL := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", cfg.User, cfg.Password, cfg.URL, cfg.Schema)
	db, err := sql.Open("mysql", schemaURL)
	if err != nil {
		logger.Error(" - Init MySQL Database Connection Failed.", err.Error())
		return nil, err
	}

	return &dbMysql{db: db}, nil
}

// VerifyMysqlConnection Verifies a connection to the database is still alive.
func (m *dbMysql) VerifyConnection() bool {
	log.Println("verify connection")
	if err := m.db.Ping(); err != nil {
		logger.Error("Failed verifies a connection to the database", err)
		return false
	}
	log.Println("success verify")
	return true
}
