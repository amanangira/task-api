package app

import (
	"fmt"
	"os"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var sqlxClient *sqlx.DB
var sqlxInit = new(sync.Once)

func GetDBDSN(
	username,
	password,
	hostname,
	port,
	database string,
	sslMode bool,
) string {
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s",
		username,
		password,
		hostname,
		port,
	)

	if database != "" {
		dsn = fmt.Sprintf("%s dbname=%s", dsn, database)
	}

	if !sslMode {
		dsn = fmt.Sprintf("%s sslmode=disable", dsn)
	}

	return dsn
}

func GetDefaultDBDSN() string {
	return GetDBDSN(
		os.Getenv(EnvDBUsernameKey),
		os.Getenv(EnvDBPasswordKey),
		os.Getenv(EnvDBHostKey),
		os.Getenv(EnvDBPortKey),
		os.Getenv(EnvDBNameKey),
		false)
}

func NewSQLXDbClientFromDSN(dsn string) *sqlx.DB {
	var sqlDB *sqlx.DB
	var err error
	sqlDB, err = sqlx.Open("postgres", dsn)
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	return sqlDB
}

func GetSQLDXBClient() *sqlx.DB {
	sqlxInit.Do(func() {
		dsn := GetDefaultDBDSN()
		sqlxClient = NewSQLXDbClientFromDSN(dsn)
	})

	return sqlxClient

}
