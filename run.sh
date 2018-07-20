#!/bin/sh

/app/sql-migrate up
/app/sql-migrate up -env=seed

/app/core