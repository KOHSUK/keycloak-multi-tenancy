package tenant

import (
	"context"

	"github.com/google/uuid"
)

type ProvisionTenantCommand struct {
	TenantId uuid.UUID
}

type ProvisionTenantUseCase interface {
	Handle(ctx context.Context, command ProvisionTenantCommand) error
}
