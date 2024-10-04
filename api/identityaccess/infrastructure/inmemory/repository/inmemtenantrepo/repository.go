package inmemtenantrepo

import (
	"api/identityaccess/domain/identity/model"
	"context"

	"github.com/google/uuid"
)

type InMemTenantRepository struct{}

func NewInMemTenantRepository() *InMemTenantRepository {
	return &InMemTenantRepository{}
}

func (r *InMemTenantRepository) TenantOfId(ctx context.Context, id model.TenantId) (*model.Tenant, error) {
	return &model.Tenant{
		ID:   id,
		Name: "Tenant 1",
	}, nil
}

func (r *InMemTenantRepository) Save(ctx context.Context, tenant *model.Tenant) error {
	return nil
}

func (r *InMemTenantRepository) NextIdentity(ctx context.Context) *model.TenantId {
	return &model.TenantId{Value: uuid.New()}
}
