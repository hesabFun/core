package companypb

import (
	"context"
	"github.com/pkg/errors"
)

// RegisterUser is to register new user
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
