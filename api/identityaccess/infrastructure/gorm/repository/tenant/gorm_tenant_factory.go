package tenant

import (
	"api/identityaccess/domain/identity/model"
	"api/identityaccess/infrastructure/gorm/connector"

	"github.com/google/uuid"
)

type GormTenantFactory struct {
	connector connector.Connector
}

func NewGormTenantFactory() *GormTenantFactory {
	return &GormTenantFactory{}
}

func (f *GormTenantFactory) NewTenant(id uuid.UUID, name string) *model.Tenant {
	// insert tenant into database and return tenant
	tenant := &model.Tenant{
		ID:     model.NewTenantId(id),
		Name:   name,
		Active: false,
	}

	db := f.connector.GetDB()
	db.Create(tenant)

	return tenant
}
