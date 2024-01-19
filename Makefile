migrate:
	dbmate migrate

migrate-down:
	dbmate down

run:
	go build -o ./.bin/server . && ./.bin/server --dev