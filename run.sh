#!/bin/sh

./sql-migrate up
./sql-migrate up -env=seed

./core