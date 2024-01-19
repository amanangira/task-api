package app

import (
	"github.com/jmoiron/sqlx"
	"log"
)

type IAPI interface {
	IsDebug() bool
	GetLogger() *log.Logger
	GetDBClient() *sqlx.DB
}
