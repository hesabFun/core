package impl

import (
	"context"
	companypb "elbix.dev/engine/modules/company/proto"
	"elbix.dev/engine/pkg/grpcgw"
	"elbix.dev/engine/pkg/mockery"
	"github.com/fullstorydev/grpchan/inprocgrpc"
	"github.com/stretchr/testify/assert"
	"testing"
)

var ch *inprocgrpc.Channel

func newClient() companypb.CompanySystemClient {
	if ch == nil {
		ch = grpcgw.GRPCChannel()
		companypb.RegisterHandlerCompanySystem(ch, companypb.NewWrappedCompanySystemServer(&companyController{}))
	}
	return companypb.NewCompanySystemChannelClient(ch)
}

func TestCompanyController_createCompany(t *testing.T) {
	ctx := context.Background()
	defer mockery.Start(ctx, t)()

	c := newClient()
	r, err := c.CreateCompany(ctx, &companypb.CreateCompanyRequest{
		Name: "test company",
	})

	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, "test company", r.GetName())
	assert.Equal(t, companypb.CompanyStatus_COMPANY_STATUS_ACTIVE, r.GetStatus())
	assert.NotNil(t, r.GetId())

	r2, err := c.GetCompany(ctx, &companypb.GetCompanyRequest{
		Id: r.GetId(),
	})
	assert.NoError(t, err)
	assert.NotNil(t, r2)
	assert.Equal(t, r.GetName(), r2.GetName())
	assert.Equal(t, r.GetStatus(), r2.GetStatus())
	assert.NotNil(t, r2.Id)
	assert.Equal(t, r.GetId(), r2.GetId())
}
