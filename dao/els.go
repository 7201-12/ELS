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

// by type and complexity
func (e *Els) GetQuestions(t models.TestType) ([]*models.Question, error) {
	questions := make([]*models.Question, 0)
	rows, err := e.DB.Query(ctx, `SELECT id, q_type, t_type, value, variants, answer, time FROM questions WHERE t_type = $1`, t)
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
		q.Answer, err = e.GetVariant(q.Answer.ID)
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

func (e *Els) GetQuestion(id uuid.UUID) (*models.Question, error) {
	question := &models.Question{}
	row := e.DB.QueryRow(ctx, `SELECT id, q_type, t_type, value, variants, answer, time FROM questions WHERE id = $1`, id)
	err := scanQuestion(row, question)
	if err != nil {
		return nil, err
	}
	return question, nil
}

func (e *Els) GetVariant(id uuid.UUID) (models.Variant, error) {
	variant := &models.Variant{}
	row := e.DB.QueryRow(ctx, `SELECT id, value FROM variants WHERE id = $1`, id)
	err := scanVariant(row, variant)
	if err != nil {
		return models.Variant{}, err
	}
	return *variant, nil
}

func scanQuestion(row pgx.Row, question *models.Question) error {
	sl := []string{}
	err := row.Scan(&question.ID, &question.Type, &question.TestType, &question.Value, pq.Array(&sl), &question.Answer.ID, &question.Time)
	if err != nil {
		return err
	}
	for _, v := range sl {
		id, err := uuid.Parse(v)
		if err != nil {
			return err
		}
		question.Variants = append(question.Variants, &models.Variant{ID: id})
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
