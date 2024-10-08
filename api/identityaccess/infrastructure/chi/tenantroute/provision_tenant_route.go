package tenantroute

import (
	"api/identityaccess/interface/controller/tenantctl"
	"encoding/json"
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

type ProvisionTenantRequest struct {
	Name string `json:"name"`
}

func (r *ProvisionTenantRoute) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var ptr ProvisionTenantRequest
	err := json.NewDecoder(req.Body).Decode(&ptr)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = r.controller.Handle(req.Context(), tenantctl.ProvisionTenantRequest{
		TenantName: ptr.Name,
	})

	if err != nil {
		r.logger.Error("Failed to provision tenant", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("OK"))
}
