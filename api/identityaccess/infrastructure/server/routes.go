package server

import "api/identityaccess/infrastructure/chi/tenantroute"

type Routes struct {
	provisionTenantRoute tenantroute.ProvisionTenantRoute
}

func NewRoutes(provisionTenantRoute *tenantroute.ProvisionTenantRoute) *Routes {
	return &Routes{
		provisionTenantRoute: *provisionTenantRoute,
	}
}
