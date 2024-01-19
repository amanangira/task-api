package app

import (
	"context"
	"github.com/jmoiron/sqlx"
	"log"
)

type API struct {
	debug  bool
	logger *log.Logger
}

func (a API) IsDebug() bool {
	return a.debug
}

func (a API) GetDBClient() *sqlx.DB {
	return GetSQLDXBClient()
}

func (a API) GetLogger() *log.Logger {
	return a.logger
}

func NewAPI(
	debug bool,
	logger *log.Logger) *API {

	return &API{
		debug:  debug,
		logger: logger,
	}
}

func InitializeAPI(ctx context.Context) *API {
	logger := log.Default()

	isDebug := IsDebug()

	return NewAPI(
		isDebug,
		logger,
	)
}
