// +build wireinject

package company

import (
	"elbix.dev/engine/modules/company/impl"
	// migrations
	_ "elbix.dev/engine/modules/company/migrations"
	companypb "elbix.dev/engine/modules/company/proto"
	"elbix.dev/engine/pkg/grpcgw"
	"elbix.dev/engine/pkg/sec"
	"github.com/google/wire"
)

// CompanySet is the builder used to build this module
var CompanySet = wire.NewSet(
	wire.Bind(new(grpcgw.Controller), new(companypb.WrappedCompanySystemController)),
	sec.ParseRSAPrivateKeyFromBase64PEM,
	sec.ExtractPublicFromPrivate,
	companypb.NewWrappedCompanySystemServer,
	impl.NewCompanyController,
)
