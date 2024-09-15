package repository

import (
	"api/identityaccess/domain/model/identity"
	"context"
)

type UserRepository interface {
	UserOfId(ctx context.Context, id identity.UserId) (*identity.User, error)
	Save(ctx context.Context, user *identity.User) error
}
