package models

import "github.com/google/uuid"

type TestType int32

// полнота целостность проблемы
const (
	Fullness TestType = iota + 1
	Integrity
	Problems
)

type ThemeType int32

const (
	Poly1 ThemeType = iota + 1
	Linear
	Poly2
)

type Question struct {
	ID       uuid.UUID  `json:"id"`
	TestId   TestType   `json:"test_id"`
	ThemeId  ThemeType  `json:"theme_id"`
	Value    string     `json:"value"`
	Variants []*Variant `json:"variants"`
	Answer   Variant    `json:"answer"`
	Time     float64    `json:"time"`
}
