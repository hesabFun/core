#!/bin/sh

cd /go/src/github.com/hesabFun/core
sql-migrate up
sql-migrate up -env=seed
./core
