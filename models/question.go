package models

import "github.com/google/uuid"

type QuestionType int32

const (
	Theory QuestionType = iota
	Methodology
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
