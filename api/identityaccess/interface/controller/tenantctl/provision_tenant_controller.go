package tenantctl

import (
	"api/identityaccess/usecase/command/tenant"
	"context"
)

type ProvisionTenantController struct {
	usecase tenant.ProvisionTenantUseCase
}

func NewProvisionTenantController(usecase tenant.ProvisionTenantUseCase) *ProvisionTenantController {
	return &ProvisionTenantController{usecase: usecase}
}

type ProvisionTenantRequest struct {
	TenantName string
}

func (c *ProvisionTenantController) Handle(ctx context.Context, req ProvisionTenantRequest) {
	if c.usecase == nil {
		panic("usecase is not initialized")
	}

	c.usecase.Handle(ctx, tenant.ProvisionTenantCommand{
		Name: req.TenantName,
	})
}
