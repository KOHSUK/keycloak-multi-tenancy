package identity

type User struct {
	ID UserId
	Person
}

func NewUser(id UserId, name FullName, tenantId TenantId) *User {
	return &User{
		ID:     id,
		Person: Person{Name: name, TenantId: tenantId},
	}
}
