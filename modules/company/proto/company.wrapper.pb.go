// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: modules/company/proto/company.proto

package companypb

import (
	elbix_dev_engine_pkg_assert "elbix.dev/engine/pkg/assert"
	elbix_dev_engine_pkg_grpcgw "elbix.dev/engine/pkg/grpcgw"
	elbix_dev_engine_pkg_resources "elbix.dev/engine/pkg/resources"
	fmt "fmt"
	_ "github.com/fzerorubigd/protobuf/extra"
	_ "github.com/fzerorubigd/protobuf/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	github_com_grpc_ecosystem_grpc_gateway_runtime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	golang_org_x_net_context "golang.org/x/net/context"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	google_golang_org_grpc "google.golang.org/grpc"
	gopkg_in_go_playground_validator_v9 "gopkg.in/go-playground/validator.v9"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type WrappedCompanySystemController interface {
	CompanySystemServer
	elbix_dev_engine_pkg_grpcgw.Controller
}

type wrappedCompanySystemServer struct {
	original CompanySystemServer
	v        *gopkg_in_go_playground_validator_v9.Validate
}

func (w *wrappedCompanySystemServer) Init(ctx golang_org_x_net_context.Context, conn *google_golang_org_grpc.ClientConn, mux *github_com_grpc_ecosystem_grpc_gateway_runtime.ServeMux) {
	cl := NewCompanySystemClient(conn)

	elbix_dev_engine_pkg_assert.Nil(RegisterCompanySystemHandlerClient(ctx, mux, cl))
}

func (w *wrappedCompanySystemServer) InitGRPC(ctx golang_org_x_net_context.Context, s *google_golang_org_grpc.Server) {
	RegisterCompanySystemServer(s, w)
}

func (w *wrappedCompanySystemServer) GetCompany(ctx golang_org_x_net_context.Context, req *GetCompanyRequest) (res *CompanyResponse, err error) {
	ctx, err = elbix_dev_engine_pkg_grpcgw.ExecuteMiddleware(ctx, w.original)
	if err != nil {
		return nil, err
	}
	if err = w.v.StructCtx(ctx, req); err != nil {
		return nil, elbix_dev_engine_pkg_grpcgw.NewBadRequest(err, "validation failed")
	}

	res, err = w.original.GetCompany(ctx, req)
	return
}

func (w *wrappedCompanySystemServer) CreateCompany(ctx golang_org_x_net_context.Context, req *CreateCompanyRequest) (res *CompanyResponse, err error) {
	ctx, err = elbix_dev_engine_pkg_grpcgw.ExecuteMiddleware(ctx, w.original)
	if err != nil {
		return nil, err
	}
	if err = w.v.StructCtx(ctx, req); err != nil {
		return nil, elbix_dev_engine_pkg_grpcgw.NewBadRequest(err, "validation failed")
	}

	res, err = w.original.CreateCompany(ctx, req)
	return
}

func (w *wrappedCompanySystemServer) GetCompanies(ctx golang_org_x_net_context.Context, req *GetCompaniesRequest) (res *CompaniesResponse, err error) {
	ctx, err = elbix_dev_engine_pkg_grpcgw.ExecuteMiddleware(ctx, w.original)
	if err != nil {
		return nil, err
	}
	if err = w.v.StructCtx(ctx, req); err != nil {
		return nil, elbix_dev_engine_pkg_grpcgw.NewBadRequest(err, "validation failed")
	}

	res, err = w.original.GetCompanies(ctx, req)
	return
}

func NewWrappedCompanySystemServer(server CompanySystemServer) WrappedCompanySystemController {
	return &wrappedCompanySystemServer{
		original: server,
		v:        gopkg_in_go_playground_validator_v9.New(),
	}
}
func init() {
	elbix_dev_engine_pkg_resources.RegisterResource("/company.CompanySystem/GetCompany", "")
	elbix_dev_engine_pkg_resources.RegisterResource("/company.CompanySystem/CreateCompany", "")
	elbix_dev_engine_pkg_resources.RegisterResource("/company.CompanySystem/GetCompanies", "")
}
