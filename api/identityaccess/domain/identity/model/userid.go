package model

import "github.com/google/uuid"

type UserId struct {
	Value uuid.UUID
}

func (u UserId) String() string {
	return u.Value.String()
}

func NewUserId(value uuid.UUID) UserId {
	return UserId{
		Value: value,
	}
}
