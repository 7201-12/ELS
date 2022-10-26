package models

import "github.com/google/uuid"

type QuestionType int32

// полнота целостность проблемы
const (
	Fulltegrity QuestionType = iota + 1
	Problems
)

type Question struct {
	ID       uuid.UUID
	Type     QuestionType
	Value    string
	Variants []*Variant
	Answer   Variant
	Time     float64
}
