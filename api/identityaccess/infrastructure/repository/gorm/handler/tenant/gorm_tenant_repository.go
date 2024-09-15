package tenant_repository

import (
	"api/identityaccess/domain/model/identity"
	"api/identityaccess/infrastructure/repository/gorm/connector"
	"api/identityaccess/infrastructure/repository/gorm/model"
	"context"

	"gorm.io/gorm/clause"
)

type GormTenantRepository struct {
	connector connector.Connector
}

func NewGormTenantRepository(connector connector.Connector) *GormTenantRepository {
	return &GormTenantRepository{connector: connector}
}

func (r *GormTenantRepository) TenantOfId(ctx context.Context, id string) (*identity.Tenant, error) {
	db := r.connector.GetDB()
	var tenant model.Tenant
	result := db.First(&tenant, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &identity.Tenant{
		ID:   identity.NewTenantId(tenant.ID),
		Name: tenant.Name,
	}, nil
}

func (r *GormTenantRepository) Save(ctx context.Context, tenant *identity.Tenant) error {
	db := r.connector.GetDB()
	result := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"name": tenant.Name}),
	}).Create(&model.Tenant{
		ID:   tenant.ID.Value,
		Name: tenant.Name,
	})
	if result.Error != nil {
		return result.Error
	}

	var users []model.User
	for _, user := range tenant.TenantMembers {
		users = append(users, model.User{
			ID:       user.UserId.Value,
			TenantId: tenant.ID.Value,
		})
	}

	result = db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"tenant_id"}),
	}).Create(&users)

	if result.Error != nil {
		return result.Error
	}

	return result.Error
}
