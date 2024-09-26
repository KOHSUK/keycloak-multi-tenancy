package tenant

import (
	"api/identityaccess/domain/identity/factory"
	"api/identityaccess/domain/identity/repository"
	"api/identityaccess/usecase/command/tenant"
	"context"
	"errors"
)

type ProvisionTenantCommandHandler struct {
	tenantRepository repository.TenantRepository
	tenantFactory    factory.TenantFactory
}

func NewProvisionTenantCommandHandler(tenantRepository *repository.TenantRepository, tenantFactory *factory.TenantFactory) *ProvisionTenantCommandHandler {
	return &ProvisionTenantCommandHandler{tenantRepository: *tenantRepository, tenantFactory: *tenantFactory}
}

func (h *ProvisionTenantCommandHandler) Handle(ctx context.Context, command tenant.ProvisionTenantCommand) error {
	tenantId, err := h.tenantRepository.NextIdentity(ctx)
	if err != nil {
		return err
	}

	tenant, err := h.tenantFactory.NewTenant(*tenantId, command.Name)
	if err != nil {
		return err
	}

	if tenant == nil {
		return errors.New("could not create tenant")
	}

	return nil
}
