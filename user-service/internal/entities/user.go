package entities

import "time"

type UserEntity struct {
	ID     uint64
	Name   string
	Email  string
	Status string

	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
}
