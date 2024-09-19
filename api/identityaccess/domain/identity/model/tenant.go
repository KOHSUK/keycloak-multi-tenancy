package model

import "errors"

type Tenant struct {
	ID            TenantId
	Name          string
	TenantMembers []TenantMember
	Active        bool
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

func (t *Tenant) Deactivate() {
	t.Active = false
}

func (t *Tenant) Activate() {
	t.Active = true
}
