build:
	go build -o core

format:
	go fmt

test: db-reset
	go test

db-remove:
	mysql -u ${MYSQL_USERNAME} -e "DROP DATABASE IF EXISTS ${MYSQL_DATABASE}"

db-create:
	mysql -u ${MYSQL_USERNAME} -e "CREATE DATABASE ${MYSQL_DATABASE}"

db-reset: db-remove db-create migrate-all

migrate:
	sql-migrate up

seed:
	sql-migrate up -env=seed

seed-test:
	sql-migrate up -env=test

migrate-all: migrate seed seed-test

doc:
	apidoc -c apidoc/config/ -o apidoc/build