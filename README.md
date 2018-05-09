# HeasbFun - core
[![Build Status](https://circleci.com/gh/hesabFun/core.svg?&style=shield)](https://circleci.com/gh/hesabFun/core)
[![Docker Repository on Quay](https://quay.io/repository/hesabfun/core/status "Docker Repository on Quay")](https://quay.io/repository/hesabfun/core)
[![codecov](https://codecov.io/gh/hesabFun/core/branch/master/graph/badge.svg)](https://codecov.io/gh/hesabFun/core)
[![Go Report](https://goreportcard.com/badge/github.com/hesabFun/core)](https://goreportcard.com/report/github.com/hesabFun/core)
[![License](https://img.shields.io/badge/License-AGPL%203.0-blue.svg)](https://github.com/hesabFun/core/blob/master/LICENSE)

Free, open source and cross-platform finance application

## After clone

First install mariadb

Set ENV, `MYSQL_DATABASE=your_db_name`, `MYSQL_ADDRESS=localhost`, `MYSQL_USERNAME=root`, `MYSQL_PASSWORD=your_password`, `MYSQL_PORT=3306`

Install sql-migrate, run `go get -v github.com/rubenv/sql-migrate/...`

Migrate tables, run `sql-migrate up`

Database seeding, run `sql-migrate up -env=seed`

Database seed test data, run `sql-migrate up -env=test`

Run `cp .env.example .env` and insert your env

Install glide, run `curl https://glide.sh/get | sh`

Install packages, run `glide install` 

## Build

Run `go build` and `./core`

## Running tests

Run `go test`

## Code scaffolding

Run `sql-migrate your_mgirate_title` to generate a new migrate file in `./migrations`.

Run `sql-migrate -env=seed your_seed_title` to generate a new seed file in `./seeds`.

Run `sql-migrate -env=test your_test_title` to generate a new seed file in `./seed_test`.
