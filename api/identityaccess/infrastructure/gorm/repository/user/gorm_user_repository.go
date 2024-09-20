package user

import (
	"api/identityaccess/domain/identity/model"
	"api/identityaccess/infrastructure/gorm/connector"
	"api/identityaccess/infrastructure/gorm/repository/person"
	"context"

	"gorm.io/gorm/clause"
)

type GormUserRepository struct {
	connector connector.Connector
}

func (r *GormUserRepository) UserOfId(ctx context.Context, id string) (*model.User, error) {
	db := r.connector.GetDB()
	var user User
	result := db.First(&user, "id = ?", id)
	if result.Error == nil {
		return nil, result.Error
	}
	return &model.User{
		ID: model.NewUserId(user.ID),
		Person: model.Person{
			ID:       model.NewUserId(user.ID),
			TenantId: model.NewTenantId(user.TenantId),
			Name:     model.NewFullName(user.Person.FirstName, user.Person.LastName),
			Email:    user.Person.Email,
		},
	}, nil
}

func (r *GormUserRepository) Save(ctx context.Context, user *model.User) error {
	db := r.connector.GetDB()
	result := db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"active": user.Active}),
	}).Create(&User{
		ID:     user.ID.Value,
		Active: user.Active,
	})

	if result.Error != nil {
		return result.Error
	}

	result = db.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"first_name": user.Person.Name.FirstName,
			"last_name":  user.Person.Name.LastName,
			"email":      user.Person.Email,
			"active":     user.Active,
		}),
	}).Create(&person.Person{
		ID: user.ID.Value,
	})

	if result.Error != nil {
		return result.Error
	}

	return nil
}
