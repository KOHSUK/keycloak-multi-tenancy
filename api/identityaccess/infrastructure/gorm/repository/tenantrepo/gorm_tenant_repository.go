package tenantrepo

import (
	"api/identityaccess/domain/identity/model"
	"api/identityaccess/infrastructure/gorm/connector"
	userModel "api/identityaccess/infrastructure/gorm/repository/user"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

type GormTenantRepository struct {
	connector connector.Connector
}

func NewGormTenantRepository(connector connector.Connector) *GormTenantRepository {
	return &GormTenantRepository{connector: connector}
}

func (r *GormTenantRepository) TenantOfId(ctx context.Context, id model.TenantId) (*model.Tenant, error) {
	db := r.connector.GetDB()
	var tenant Tenant
	result := db.First(&tenant, "id = ?", id.Value)
	if result.Error != nil {
		return nil, result.Error
	}
	return &model.Tenant{
		ID:   model.NewTenantId(tenant.ID),
		Name: tenant.Name,
	}, nil
}

func (r *GormTenantRepository) Save(ctx context.Context, tenant *model.Tenant) error {
	db := r.connector.GetDB()
	result := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"name": tenant.Name}),
	}).Create(&Tenant{
		ID:     tenant.ID.Value,
		Name:   tenant.Name,
		Active: tenant.Active,
	})
	if result.Error != nil {
		return result.Error
	}

	var users []userModel.User
	for _, user := range tenant.TenantMembers {
		users = append(users, userModel.User{
			ID:       user.UserId.Value,
			TenantId: tenant.ID.Value,
			Active:   tenant.Active,
		})
	}

	// Save Users
	result = db.Clauses(clause.OnConflict{DoNothing: true}).Create(&users)

	if result.Error != nil {
		return result.Error
	}

	return result.Error
}

func (r *GormTenantRepository) NextIdentity(ctx context.Context) *model.TenantId {
	return &model.TenantId{Value: uuid.New()}
}
