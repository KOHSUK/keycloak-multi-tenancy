package tenantroute

import (
	"api/identityaccess/interface/controller/tenantctl"
	"net/http"
)

type ProvisionTenantRoute struct {
	controller tenantctl.ProvisionTenantController
}

func NewProvisionTenantRoute(controller *tenantctl.ProvisionTenantController) *ProvisionTenantRoute {
	return &ProvisionTenantRoute{
		controller: *controller,
	}
}

func (r *ProvisionTenantRoute) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.controller.Handle(req.Context(), tenantctl.ProvisionTenantRequest{
		TenantName: req.FormValue("tenant_name"),
	})

	w.Write([]byte("OK"))
}
