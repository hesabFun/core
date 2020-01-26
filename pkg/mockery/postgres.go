package mockery

import (
	"context"
	"database/sql"

	"elbix.dev/engine/pkg/assert"
	"elbix.dev/engine/pkg/config"
	"elbix.dev/engine/pkg/random"
)

var (
	user    = config.RegisterString("mockery.postgres.user", "engine_test", "postgres user")
	dbname  = config.RegisterString("mockery.postgres.db", "engine_test", "postgres database")
	pass    = config.RegisterString("mockery.postgres.password", "bita123", "postgres password")
	host    = config.RegisterString("mockery.postgres.host", "localhost", "postgres host")
	port    = config.RegisterInt("mockery.postgres.port", 5432, "postgres port")
	sslmode = config.RegisterString("mockery.postgres.sslmode", "disable", "sslmode for postgres")
)

func sqltxTesting(ctx context.Context) (*sql.DB, error) {
	db, err := sql.Open("txdb", <-random.ID)
	assert.Nil(err)
	go func() {
		<-ctx.Done()
		assert.Nil(db.Close())
	}()
	return db, nil
}
