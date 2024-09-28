package tenantctl

import (
	"api/identityaccess/usecase/command/uctenant"
	"context"
)

type ProvisionTenantController struct {
	usecase uctenant.ProvisionTenantUseCase
}

func NewProvisionTenantController(usecase uctenant.ProvisionTenantUseCase) *ProvisionTenantController {
	return &ProvisionTenantController{usecase: usecase}
}

type ProvisionTenantRequest struct {
	TenantName string
}

func (c *ProvisionTenantController) Handle(ctx context.Context, req ProvisionTenantRequest) {
	if c.usecase == nil {
		panic("usecase is not initialized")
	}

	c.usecase.Handle(ctx, uctenant.ProvisionTenantCommand{
		Name: req.TenantName,
	})
}
