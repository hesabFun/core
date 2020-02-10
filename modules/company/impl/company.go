package impl

import (
	"context"
	companypb "elbix.dev/engine/modules/company/proto"
	"elbix.dev/engine/pkg/assert"
	"elbix.dev/engine/pkg/grpcgw"
)

type companyController struct{}

func (uc *companyController) CreateCompany(ctx context.Context, rc *companypb.CreateCompanyRequest) (*companypb.CreateCompanyResponse, error) {
	m := companypb.NewManager()

	c, err := m.AddCompany(ctx, rc.GetName())
	if err != nil {
		return nil, grpcgw.NewBadRequest(err, "company could not insert")
	}

	return &companypb.CreateCompanyResponse{
		Id:     c.GetId(),
		Status: c.GetStatus(),
		Name:   c.GetName(),
	}, nil
}

func (uc *companyController) GetCompany(ctx context.Context, rc *companypb.GetCompanyRequest) (*companypb.GetCompanyResponse, error) {
	m := companypb.NewManager()

	c, err := m.GetCompanyByPrimary(ctx, rc.GetId())
	assert.Nil(err)

	return &companypb.GetCompanyResponse{
		Id:     c.GetId(),
		Name:   c.GetName(),
		Status: c.GetStatus(),
	}, nil
}

// NewCompanyController return a grpc user controller
func NewCompanyController() companypb.CompanyServiceServer {
	return &companyController{}
}
