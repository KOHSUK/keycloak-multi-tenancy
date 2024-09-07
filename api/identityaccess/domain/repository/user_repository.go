package repository

import (
	"api/identityaccess/domain/model/identity"
	"context"
)

type UserRepository interface {
	UserOfId(ctx context.Context, id string) (*identity.User, error)
	Add(ctx context.Context, user *identity.User) error
}
