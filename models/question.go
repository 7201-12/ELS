package models

import "github.com/google/uuid"

type QuestionType int32

// полнота целостность проблемы
const (
	Fullness QuestionType = iota + 1
	Integrity
	Problems
)

type TestType int32

const (
	Poly1 TestType = iota + 1
	Linear
	Poly2
)

type Question struct {
	ID       uuid.UUID    `json:"id"`
	Type     QuestionType `json:"type"`
	TestType TestType     `json:"t_type"`
	Value    string       `json:"value"`
	Variants []*Variant   `json:"variants"`
	Answer   Variant      `json:"answer"`
	Time     float64      `json:"time"`
}
