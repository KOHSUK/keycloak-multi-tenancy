package tenant

import (
	"api/identityaccess/domain/identity/model"
	"api/identityaccess/domain/identity/repository"
	"api/identityaccess/usecase/command/tenant"
	"context"
	"errors"
)

type ProvisionTenantCommandHandler struct {
	tenantRepository repository.TenantRepository
}

func NewProvisionTenantCommandHandler(tenantRepository repository.TenantRepository) *ProvisionTenantCommandHandler {
	return &ProvisionTenantCommandHandler{tenantRepository: tenantRepository}
}

func (h *ProvisionTenantCommandHandler) Handle(ctx context.Context, command tenant.ProvisionTenantCommand) error {
	tenant, err := h.tenantRepository.TenantOfId(ctx, model.NewTenantId(command.TenantId))
	if err != nil {
		return err
	}

	if tenant != nil && tenant.Active {
		return errors.New("tenant already provisioned")
	}

	return nil
}
