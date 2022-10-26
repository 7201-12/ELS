package dao

import (
	"context"

	"github.com/7201-12/ELS/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lib/pq"
)

var ctx = context.Background()

type Els struct {
	DB *pgxpool.Pool
}

func (e *Els) GetQuestionsByType(t int32) ([]*models.Question, error) {
	questions := make([]*models.Question, 0)
	rows, err := e.DB.Query(ctx, `SELECT id, q_type, value, variants, answer, time FROM questions WHERE q_type = $1`, t)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		q := &models.Question{}
		err = scanQuestion(rows, q)
		if err != nil {
			return nil, err
		}
		for _, v := range q.Variants {
			variant, err := e.GetVariant(v.ID)
			if err != nil {
				return nil, err
			}
			v.Value = variant.Value
		}
		questions = append(questions, q)
	}
	return questions, nil
}

func (e *Els) GetVariant(id uuid.UUID) (*models.Variant, error) {
	variant := &models.Variant{}
	row := e.DB.QueryRow(ctx, `SELECT id, value FROM variants WHERE id = $1`, id)
	err := scanVariant(row, variant)
	if err != nil {
		return nil, err
	}
	return variant, nil
}

func scanQuestion(row pgx.Row, question *models.Question) error {
	err := row.Scan(&question.ID, &question.Type, &question.Value, pq.Array(&question.Variants), &question.Answer)
	if err != nil {
		return err
	}
	return nil
}

func scanVariant(row pgx.Row, variant *models.Variant) error {
	err := row.Scan(&variant.ID, &variant.Value)
	if err != nil {
		return err
	}
	return nil
}
