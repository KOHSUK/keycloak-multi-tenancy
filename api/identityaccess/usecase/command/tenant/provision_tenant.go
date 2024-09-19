package tenant

import "context"

type ProvisionTenantCommand struct {
	TenantId string
}

type ProvisionTenantUseCase interface {
	Handle(ctx context.Context, command ProvisionTenantCommand) error
}
