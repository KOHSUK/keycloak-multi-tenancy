package identity

import "errors"

type Tenant struct {
	ID            TenantId
	Name          string
	TenantMembers []TenantMember
}

func (t *Tenant) AddUser(user User) error {
	if t.ID != user.TenantId {
		return errors.New("")
	}

	member := TenantMember{
		TenantId: t.ID,
		UserId:   user.ID,
	}
	t.TenantMembers = append(t.TenantMembers, member)
	return nil
}
