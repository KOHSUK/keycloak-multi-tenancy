package tenant_repository

import (
	"api/identityaccess/domain/model/identity"
	"api/identityaccess/infrastructure/datasource/gorm/connector"
	"api/identityaccess/infrastructure/datasource/gorm/model/tenant_model"
	"context"
)

type GormTenantRepository struct {
	connector connector.Connector
}

func NewGormTenantRepository(connector connector.Connector) *GormTenantRepository {
	return &GormTenantRepository{connector: connector}
}

func (r *GormTenantRepository) GetTenantById(ctx context.Context, id string) (*identity.Tenant, error) {
	db := r.connector.GetDB()
	var tenant tenant_model.Tenant
	result := db.First(&tenant, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &identity.Tenant{
		ID:   identity.NewTenantId(tenant.ID),
		Name: tenant.Name,
	}, nil
}
