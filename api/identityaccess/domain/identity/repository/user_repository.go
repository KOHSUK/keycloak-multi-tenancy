package repository

import (
	"api/identityaccess/domain/identity/model"
	"context"
)

type UserRepository interface {
	UserOfId(ctx context.Context, id model.UserId) (*model.User, error)
	Save(ctx context.Context, user *model.User) error
}
