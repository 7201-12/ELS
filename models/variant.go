package models

import "github.com/google/uuid"

type Variant struct {
	ID    uuid.UUID `json:"id"`
	Value string    `json:"value"`
}
