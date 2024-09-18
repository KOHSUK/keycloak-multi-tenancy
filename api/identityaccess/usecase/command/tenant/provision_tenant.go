package tenant

import "context"

type ProvisionTenantCommand struct {

}

type ProvisionTenantUseCase interface {
	Handle(ctx context.Context, command ProvisionTenantCommand) error
}