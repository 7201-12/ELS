package models

import "github.com/google/uuid"

type Variant struct {
	ID    uuid.UUID
	Value string
}
