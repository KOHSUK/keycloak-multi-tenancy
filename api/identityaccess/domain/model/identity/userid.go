package identity

import "github.com/google/uuid"

type UserId struct {
	Value uuid.UUID
}

func (u UserId) String() string {
	return u.Value.String()
}
