package tenant

import (
	"api/identityaccess/domain/model/identity"
	"api/identityaccess/domain/repository"
	"api/identityaccess/usecase/tenant/add_user"
	"context"
	"errors"
)

type AddUserInteractor struct {
	userRepository   repository.UserRepository
	tenantRepository repository.TenantRepository
}

func NewIdentityApplicationService(userRepository repository.UserRepository) *AddUserInteractor {
	return &AddUserInteractor{userRepository: userRepository}
}

func (ias *AddUserInteractor) AddUserToTenant(ctx context.Context, input add_user.AddUserToTenantInputData) error {
	tenant, err := ias.ExistingTenant(ctx, input.TenantId)
	if err != nil {
		return err
	}

	user, err := ias.ExistingUser(ctx, input.TenantId, input.UserId)
	if err != nil {
		return err
	}

	if err := tenant.AddUser(*user); err != nil {
		return err
	}

	return nil
}

func (ias *AddUserInteractor) User(ctx context.Context, tenantId string, userId string) (*identity.User, error) {
	user, err := ias.userRepository.UserOfId(ctx, userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ias *AddUserInteractor) ExistingUser(ctx context.Context, tenantId string, userId string) (*identity.User, error) {
	user, err := ias.User(ctx, tenantId, userId)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (ias *AddUserInteractor) Tenant(ctx context.Context, tenantId string) (*identity.Tenant, error) {
	tenant, err := ias.tenantRepository.TenantOfId(ctx, tenantId)
	if err != nil {
		return nil, err
	}

	return tenant, nil
}

func (ias *AddUserInteractor) ExistingTenant(ctx context.Context, tenantId string) (*identity.Tenant, error) {
	tenant, error := ias.Tenant(ctx, tenantId)
	if error != nil {
		return nil, error
	}

	if tenant == nil {
		return nil, errors.New("tenant not found")
	}

	return tenant, nil
}
