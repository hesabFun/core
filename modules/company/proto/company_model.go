package companypb

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"strings"
)

// AddCompany is to register new company
func (m *Manager) AddCompany(ctx context.Context, name string) (*Company, error) {
	u := Company{
		Name:   name,
		Status: CompanyStatus_COMPANY_STATUS_ACTIVE,
	}

	if err := m.CreateCompany(ctx, &u); err != nil {
		return nil, errors.Wrap(err, "could not insert")
	}

	return &u, nil
}

func (m *Manager) GetCompanies(ctx context.Context) (*CompaniesResponse, error) {
	q := fmt.Sprintf(
		"SELECT %s FROM %s WHERE id >= $1 ",
		strings.Join(m.getCompanyFields(), ","),
		CompanyTableFull,
	)

	//var companies []*Company
	var companies []*CompanyResponse
	var companiesResponse *CompaniesResponse

	//todo: get user's companies
	rows, err := m.GetDbMap().QueryxContext(ctx, q, 0)
	if err != nil {
		return companiesResponse, err
	}

	for rows.Next() {
		c, err := m.scanCompany(rows)
		if err != nil {
			return companiesResponse, err
		}
		companies = append(companies, &CompanyResponse{
			Id:     c.GetId(),
			Name:   c.GetName(),
			Status: c.GetStatus(),
		})
	}

	companiesResponse = &CompaniesResponse{Companies: companies}
	return companiesResponse, nil
}

func (m *Manager) DeleteCompany(ctx context.Context, id int64) error {
	c, err := m.GetCompanyByPrimary(ctx, id)
	if err != nil {
		return err
	}

	q := fmt.Sprintf(
		"DELETE FROM %s WHERE id = $1 ",
		CompanyTableFull,
	)
	_, err = m.GetDbMap().ExecContext(ctx, q, c.Id)

	return err
}
