package tenantrepo

import (
	"api/identityaccess/domain/identity/model"
	"api/identityaccess/infrastructure/gorm/connector"
)

type GormTenantFactory struct {
	connector connector.Connector
}

func NewGormTenantFactory() *GormTenantFactory {
	return &GormTenantFactory{}
}

func (f *GormTenantFactory) NewTenant(id model.TenantId, name string) *model.Tenant {
	// insert tenant into database and return tenant
	tenant := &model.Tenant{
		ID:     id,
		Name:   name,
		Active: false,
	}

	db := f.connector.GetDB()
	db.Create(tenant)

	return tenant
}
