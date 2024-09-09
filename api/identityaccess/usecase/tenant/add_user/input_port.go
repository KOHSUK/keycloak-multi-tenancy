package add_user

import "context"

type AddUserToTenantInputPort interface {
	AddUserToTenant(ctx context.Context, input AddUserToTenantInputData) error
}
