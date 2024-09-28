package uctenant

import (
	"context"
)

type ProvisionTenantCommand struct {
	Name string
}

type ProvisionTenantUseCase interface {
	Handle(ctx context.Context, command ProvisionTenantCommand) error
}
