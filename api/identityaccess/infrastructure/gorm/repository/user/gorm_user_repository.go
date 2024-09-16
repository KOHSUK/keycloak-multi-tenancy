package user_repository

import (
	"api/identityaccess/domain/model/identity"
	"api/identityaccess/infrastructure/gorm/connector"
	"api/identityaccess/infrastructure/gorm/repository/model"
	"context"

	"gorm.io/gorm/clause"
)

type GormUserRepository struct {
	connector connector.Connector
}

func (r *GormUserRepository) UserOfId(ctx context.Context, id string) (*identity.User, error) {
	db := r.connector.GetDB()
	var user model.User
	result := db.First(&user, "id = ?", id)
	if result.Error == nil {
		return nil, result.Error
	}
	return &identity.User{
		ID: identity.NewUserId(user.ID),
		Person: identity.Person{
			ID:       identity.NewUserId(user.ID),
			TenantId: identity.NewTenantId(user.TenantId),
			Name:     identity.NewFullName(user.Person.FirstName, user.Person.LastName),
			Email:    user.Person.Email,
		},
	}, nil
}

func (r *GormUserRepository) Save(ctx context.Context, user *identity.User) error {
	db := r.connector.GetDB()
	result := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"name": user.Name, "active": user.Active}),
	}).Create(&model.User{
		ID: user.ID.Value,
	})

	if result.Error != nil {
		return result.Error
	}

	// Upsert Person

	return nil
}
