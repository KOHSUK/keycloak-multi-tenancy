package tenantrepo

import (
	"api/identityaccess/domain/identity/model"
	"api/identityaccess/infrastructure/gorm/connector"
	"fmt"
)

type GormTenantFactory struct {
	connector connector.Connector
}

func NewGormTenantFactory(connector connector.Connector) *GormTenantFactory {
	return &GormTenantFactory{connector: connector}
}

func (f *GormTenantFactory) NewTenant(id model.TenantId, name string) (*model.Tenant, error) {
	// insert tenant into database and return tenant
	tenant := &Tenant{
		ID:     id.Value,
		Name:   name,
		Active: false,
	}

	db := f.connector.GetDB()
	err := db.Create(tenant).Error
	if err != nil {
		fmt.Println("Error creating tenant:", err)
		return nil, err
	}

	return &model.Tenant{
		ID:   id,
		Name: name,
	}, nil
}
