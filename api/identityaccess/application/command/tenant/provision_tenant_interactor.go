package tenant

import (
	"api/identityaccess/domain/identity/factory"
	"api/identityaccess/domain/identity/model"
	"api/identityaccess/domain/identity/repository"
	"api/identityaccess/usecase/command/tenant"
	"context"
	"errors"
)

type ProvisionTenantCommandHandler struct {
	tenantRepository repository.TenantRepository
	tenantFactory factory.TenantFactory
}

func NewProvisionTenantCommandHandler(tenantRepository repository.TenantRepository) *ProvisionTenantCommandHandler {
	return &ProvisionTenantCommandHandler{tenantRepository: tenantRepository}
}

func (h *ProvisionTenantCommandHandler) Handle(ctx context.Context, command tenant.ProvisionTenantCommand) error {
	tenantId, err := h.tenantRepository.NextIdentity()

	if err != nil {
		return err
	}

	tenant, err := h.tenantFactory.NewTenant(*tenantId, command.TenantId.String())

	if err != nil {
		return err
	}

	if tenant == nil {
		return errors.New("could not provision tenant")
	}

	return nil
}
