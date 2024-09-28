package tenant

import (
	"api/identityaccess/domain/identity/model"
	"api/identityaccess/domain/identity/repository"
	"api/identityaccess/usecase/command/uctenant"
	"context"
)

type AddUserToTenantCommandHandler struct {
	userRepository   repository.UserRepository
	tenantRepository repository.TenantRepository
}

func NewAddUserToTenantCommandHandler(userRepository repository.UserRepository) *AddUserToTenantCommandHandler {
	return &AddUserToTenantCommandHandler{userRepository: userRepository}
}

func (h *AddUserToTenantCommandHandler) Handle(ctx context.Context, command uctenant.AddUserToTenantCommand) error {
	// FIXME: use Factory instead of repository
	user, err := h.userRepository.UserOfId(ctx, model.NewUserId(command.UserID))
	if err != nil {
		return err
	}

	tenant, err := h.tenantRepository.TenantOfId(ctx, model.NewTenantId(command.TenantId))
	if err != nil {
		return err
	}

	tenant.AddUser(*user)

	// TODO: Save tenant

	return nil
}
