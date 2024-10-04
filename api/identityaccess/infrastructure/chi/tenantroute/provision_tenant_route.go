package tenantroute

import (
	"api/identityaccess/interface/controller/tenantctl"
	"net/http"

	"go.uber.org/zap"
)

type ProvisionTenantRoute struct {
	controller *tenantctl.ProvisionTenantController
	logger     *zap.Logger
}

func NewProvisionTenantRoute(controller *tenantctl.ProvisionTenantController, logger *zap.Logger) *ProvisionTenantRoute {
	return &ProvisionTenantRoute{
		controller: controller,
		logger:     logger,
	}
}

func (r *ProvisionTenantRoute) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := r.controller.Handle(req.Context(), tenantctl.ProvisionTenantRequest{
		TenantName: req.FormValue("tenant_name"),
	})

	if err != nil {
		r.logger.Error("Failed to provision tenant", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("OK"))
}
