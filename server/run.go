package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arymaulanamalik/bobobox_test/endpoint/routes"
	"github.com/arymaulanamalik/bobobox_test/pkg/logger"
	"github.com/arymaulanamalik/bobobox_test/repository"
	"github.com/arymaulanamalik/bobobox_test/repository/storages/mysql"
	"github.com/arymaulanamalik/bobobox_test/service"
)

type Runner struct {
	authenticateService service.AuthenticateService
	repository          repository.Repository
}

func NewRun() (*Runner, error) {
	repository, err := repository.NewRepository(repository.RepositoryConfigs{
		MysqlConfig: mysql.MysqlConfig{
			URL:      "127.0.0.1:3306",
			Schema:   "test",
			User:     "root",
			Password: "root",
		},
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	authenticateService := service.NewAuthenticateService(*repository)

	return &Runner{
		authenticateService: authenticateService,
	}, nil
}

func (r *Runner) Run(port string) error {
	router := routes.GetRouter(r.authenticateService, r.repository)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		msg := fmt.Sprintf("Failed running service without TLS. Listening on port:%s bind: address already in use", port)
		logger.Fatal(msg, nil)
	}

	return nil
}
