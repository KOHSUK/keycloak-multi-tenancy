package tenant

import (
	"context"

	"github.com/google/uuid"
)

type AddUserToTenantCommand struct {
	UserID   uuid.UUID
	Name     string
	TenantId uuid.UUID
}

type AddUserToTenantUseCase interface {
	Handle(ctx context.Context, command AddUserToTenantCommand) error
}
