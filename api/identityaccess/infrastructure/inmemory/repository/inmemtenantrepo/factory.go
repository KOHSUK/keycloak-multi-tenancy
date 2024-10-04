package inmemtenantrepo

import "api/identityaccess/domain/identity/model"

type InMemTenantFactory struct {
}

func NewInMemTenantFactory() *InMemTenantFactory {
	return &InMemTenantFactory{}
}

func (f *InMemTenantFactory) NewTenant(id model.TenantId, name string) (*model.Tenant, error) {
	// insert tenant into database and return tenant
	tenant := &model.Tenant{
		ID:     id,
		Name:   name,
		Active: false,
	}

	return tenant, nil
}
