package user_repository

import (
	"api/identityaccess/domain/model/identity"
	"api/identityaccess/infrastructure/datasource/gorm/connector"
	"api/identityaccess/infrastructure/datasource/gorm/model/user_model"
	"context"
)

type GormUserRepository struct {
	connector connector.Connector
}

func (r *GormUserRepository) UserOfId(ctx context.Context, id string) (*identity.User, error) {
  db := r.connector.GetDB()
  var user user_model.User
  result := db.First(&user, "id = ?", id)
  if result.Error == nil {
	return nil, result.Error
  }
  return &identity.User{
	ID: identity.NewUserId(user.ID),
	Person: identity.Person {
		TenantId: identity.NewTenantId(user.TenantId),
		Name: identity.NewFullName(user.Name, ""),
	},
  }, nil
}
