package models

import "github.com/google/uuid"

type QuestionType int32

// полнота целостность проблемы
const (
	Fulltegrity QuestionType = iota + 1
	Problems
)

type Question struct {
	ID       uuid.UUID    `json:"id"`
	Type     QuestionType `json:"type"`
	Value    string       `json:"value"`
	Variants []*Variant   `json:"variants"`
	Answer   Variant      `json:"answer"`
	Time     float64      `json:"time"`
}
