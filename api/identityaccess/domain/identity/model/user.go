package model

type User struct {
	ID UserId
	Person
	Active bool
}

func NewUser(id UserId, name FullName, tenantId TenantId) *User {
	return &User{
		ID:     id,
		Person: Person{Name: name, TenantId: tenantId},
		Active: false,
	}
}
