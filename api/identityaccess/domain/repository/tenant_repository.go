package repository

import (
	"api/identityaccess/domain/model/identity"
	"context"
)

type TenantRepository interface {
	TenantOfId(ctx context.Context, id identity.TenantId) (*identity.Tenant, error)
	Save(ctx context.Context, tenant *identity.Tenant) error
	NextIdentity(ctx context.Context) (*identity.TenantId, error)
}
