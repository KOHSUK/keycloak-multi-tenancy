package user

import (
	"api/identityaccess/domain/model/identity"
	"api/identityaccess/domain/repository"
	"api/identityaccess/usecase/command/tenant"
	"context"
)

type AddUserToTenantCommandHandler struct {
	userRepository   repository.UserRepository
	tenantRepository repository.TenantRepository
}

func NewAddUserToTenantCommandHandler(userRepository repository.UserRepository) *AddUserToTenantCommandHandler {
	return &AddUserToTenantCommandHandler{userRepository: userRepository}
}

func (h *AddUserToTenantCommandHandler) Handle(ctx context.Context, command tenant.AddUserToTenantCommand) error {
	user, err := h.userRepository.UserOfId(ctx, identity.NewUserId(command.UserID))
	if err != nil {
		return err
	}

	tenant, err := h.tenantRepository.TenantOfId(ctx, identity.NewTenantId(command.TenantId))
	if err != nil {
		return err
	}

	tenant.AddUser(*user)

	// TODO: Save tenant

	return nil
}
