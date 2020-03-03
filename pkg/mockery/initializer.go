package mockery

import (
	"context"
	"fmt"
	"github.com/DATA-DOG/go-txdb"
	"testing"

	"elbix.dev/engine/pkg/config"
	"elbix.dev/engine/pkg/initializer"
	"elbix.dev/engine/pkg/log"
	"elbix.dev/engine/pkg/postgres"
	"go.uber.org/zap/zaptest"
)

var (
	alreadyRegistered bool
)

// Start the mockery, used for tests only
func Start(ctx context.Context, t *testing.T) func() {
	if !alreadyRegistered {
		alreadyRegistered = true
		config.Initialize(ctx, "testing", "E")
		log.SwapLogger(zaptest.NewLogger(t))

		dsn := fmt.Sprintf(
			"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			host.String(),
			port.Int(),
			user.String(),
			pass.String(),
			dbname.String(),
			sslmode.String(),
		)
		txdb.Register("txdb", "postgres", dsn)
		postgres.DefaultInitDB = sqltxTesting
	}

	return initializer.Initialize(ctx)
}
