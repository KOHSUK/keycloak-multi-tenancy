package tenant

import "context"

type ActivateTenantCommand struct {
	TenantId string
}

type ActivateTenantUseCase interface {
	Handle(ctx context.Context, command ActivateTenantCommand) error
}
