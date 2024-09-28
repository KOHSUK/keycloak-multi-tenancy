package tenant

import (
	"api/identityaccess/domain/identity/factory"
	"api/identityaccess/domain/identity/repository"
	"api/identityaccess/usecase/command/uctenant"
	"context"
	"errors"
)

type ProvisionTenantCommandHandler struct {
	tenantRepository repository.TenantRepository
	tenantFactory    factory.TenantFactory
}

func NewProvisionTenantCommandHandler(tenantRepository repository.TenantRepository, tenantFactory factory.TenantFactory) *ProvisionTenantCommandHandler {
	return &ProvisionTenantCommandHandler{tenantRepository: tenantRepository, tenantFactory: tenantFactory}
}

func (h *ProvisionTenantCommandHandler) Handle(ctx context.Context, command uctenant.ProvisionTenantCommand) error {
	tenantId := h.tenantRepository.NextIdentity(ctx)

	tenant, err := h.tenantFactory.NewTenant(*tenantId, command.Name)
	if err != nil {
		return err
	}

	if tenant == nil {
		return errors.New("could not create tenant")
	}

	return nil
}
