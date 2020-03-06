package impl

import (
	"context"
	companypb "elbix.dev/engine/modules/company/proto"
	"elbix.dev/engine/pkg/assert"
	"elbix.dev/engine/pkg/grpcgw"
)

type companyController struct{}

func (uc *companyController) CreateCompany(ctx context.Context, rc *companypb.CreateCompanyRequest) (*companypb.CompanyResponse, error) {
	m := companypb.NewManager()

	c, err := m.AddCompany(ctx, rc.GetName())
	if err != nil {
		return nil, grpcgw.NewBadRequest(err, "company could not insert")
	}

	return &companypb.CompanyResponse{
		Id:     c.GetId(),
		Status: c.GetStatus(),
		Name:   c.GetName(),
	}, nil
}

func (uc *companyController) GetCompany(ctx context.Context, rc *companypb.GetCompanyRequest) (*companypb.CompanyResponse, error) {
	m := companypb.NewManager()

	c, err := m.GetCompanyByPrimary(ctx, rc.GetId())
	if err != nil {
		return nil, grpcgw.NewBadRequest(err, "company could not found")
	}

	return &companypb.CompanyResponse{
		Id:     c.GetId(),
		Name:   c.GetName(),
		Status: c.GetStatus(),
	}, nil
}

func (uc *companyController) GetCompanies(ctx context.Context, _ *companypb.GetCompaniesRequest) (*companypb.CompaniesResponse, error) {
	m := companypb.NewManager()

	c, err := m.GetCompanies(ctx)
	assert.Nil(err)

	return c, nil
}

func (uc *companyController) DeleteCompany(ctx context.Context, rc *companypb.DeleteCompanyRequest) (*companypb.DeleteCompanyResponse, error) {
	m := companypb.NewManager()

	_, err := m.GetCompanyByPrimary(ctx, rc.GetId())
	if err != nil {
		return nil, grpcgw.NewBadRequest(err, "company could not found")
	}

	err = m.DeleteCompany(ctx, rc.GetId())
	assert.Nil(err)

	return &companypb.DeleteCompanyResponse{}, nil
}

func (uc *companyController) UpdateCompany(ctx context.Context, rc *companypb.UpdateCompanyRequest) (*companypb.CompanyResponse, error) {
	m := companypb.NewManager()

	c, err := m.GetCompanyByPrimary(ctx, rc.GetId())
	if err != nil {
		return nil, grpcgw.NewBadRequest(err, "company could not found")
	}

	c.Name = rc.GetName()
	c.Status = rc.GetStatus()

	err = m.UpdateCompany(ctx, c)
	assert.Nil(err)

	return &companypb.CompanyResponse{}, nil
}

// NewCompanyController return a grpc user controller
func NewCompanyController() companypb.CompanySystemServer {
	return &companyController{}
}
