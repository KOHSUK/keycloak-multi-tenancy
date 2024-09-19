package model

type Person struct {
	ID       UserId
	TenantId TenantId
	Name     FullName
	Email    string
}
