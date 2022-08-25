package repository

import (
	"github.com/arymaulanamalik/bobobox_test/pkg/logger"
	"github.com/arymaulanamalik/bobobox_test/repository/storages"
	"github.com/arymaulanamalik/bobobox_test/repository/storages/mysql"
)

type RepositoryConfigs struct {
	MysqlConfig mysql.MysqlConfig
}

type Repository struct {
	ReadWriter storages.ReadWriter
}

func NewRepository(rc RepositoryConfigs) (*Repository, error) {
	logger.Info("Init Repository.", nil)
	//database
	readWriter, err := mysql.NewMysql(rc.MysqlConfig)
	if err != nil {
		return nil, err
	}

	return &Repository{
		ReadWriter: readWriter,
	}, nil
}
