package repository

import (
	"api/identityaccess/domain/identity/model"
	"context"
)

type TenantRepository interface {
	TenantOfId(ctx context.Context, id model.TenantId) (*model.Tenant, error)
	Save(ctx context.Context, tenant *model.Tenant) error
	NextIdentity(ctx context.Context) (*model.TenantId, error)
}
