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

func (c *ProvisionTenantController) Handle(ctx context.Context, req ProvisionTenantRequest) error {
	if c.usecase == nil {
		panic("usecase is not initialized")
	}

	err := c.usecase.Handle(ctx, uctenant.ProvisionTenantCommand{
		Name: req.TenantName,
	})

	if err != nil {
		return err
	}

	return nil
}
