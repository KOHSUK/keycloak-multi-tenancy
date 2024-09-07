package usecase

import (
	"api/identityaccess/domain/model/identity"
	"api/identityaccess/domain/repository"
	"api/identityaccess/usecase/command"
	"context"
	"errors"
)

type IdentityApplicationService struct {
	userRepository   repository.UserRepository
	tenantRepository repository.TenantRepository
}

func NewIdentityApplicationService(userRepository repository.UserRepository) *IdentityApplicationService {
	return &IdentityApplicationService{userRepository: userRepository}
}

func (ias *IdentityApplicationService) AddUserToTenant(ctx context.Context, cmd command.AddUserToTenantCommand) error {
	tenant, err := ias.ExistingTenant(ctx, cmd.TenantId)
	if err != nil {
		return err
	}

	user, err := ias.ExistingUser(ctx, cmd.TenantId, cmd.UserId)
	if err != nil {
		return err
	}

	if err := tenant.AddUser(*user); err != nil {
		return err
	}

	return nil
}

func (ias *IdentityApplicationService) User(ctx context.Context, tenantId string, userId string) (*identity.User, error) {
	user, err := ias.userRepository.UserOfId(ctx, userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ias *IdentityApplicationService) ExistingUser(ctx context.Context, tenantId string, userId string) (*identity.User, error) {
	user, err := ias.User(ctx, tenantId, userId)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (ias *IdentityApplicationService) Tenant(ctx context.Context, tenantId string) (*identity.Tenant, error) {
	tenant, err := ias.tenantRepository.TenantOfId(ctx, tenantId)
	if err != nil {
		return nil, err
	}

	return tenant, nil
}

func (ias *IdentityApplicationService) ExistingTenant(ctx context.Context, tenantId string) (*identity.Tenant, error) {
	tenant, error := ias.Tenant(ctx, tenantId)
	if error != nil {
		return nil, error
	}

	if tenant == nil {
		return nil, errors.New("tenant not found")
	}

	return tenant, nil
}
