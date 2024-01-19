package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/amacneil/dbmate/pkg/dbmate"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"net/url"
	"runtime/debug"
	"task/app"

	_ "github.com/amacneil/dbmate/pkg/driver/postgres"

	"os"
)

var dbmateClient *dbmate.DB

func init() {
	urlString := fmt.Sprintf(`postgres://%s:%s@%s:%s/%s`,
		os.Getenv(app.EnvDBUsernameKey),
		os.Getenv(app.EnvDBPasswordKey),
		os.Getenv(app.EnvDBHostKey),
		os.Getenv(app.EnvDBPortKey),
		os.Getenv(app.EnvDBNameKey))

	dbURL, parseErr := url.Parse(urlString)
	if parseErr != nil {
		panic(parseErr)
	}

	dbmateClient = dbmate.New(dbURL)
}

type input struct {
	Action string `json:"action"`
}

func main() {
	lambda.Start(logAndHandle)
}

func logAndHandle(ctx context.Context, i input) error {
	err := handle(ctx, i)
	if err != nil {
		log.Println(err.Error())
		log.Println(string(debug.Stack()))

		return err
	}

	return nil
}

func handle(ctx context.Context, i input) (err error) {
	switch i.Action {
	case "up", "migrate":
		err = dbmateClient.CreateAndMigrate()
		return

	case "down":
		err = dbmateClient.Rollback()
		return

	default:
		err = errors.New("invalid action")
		return
	}
}
