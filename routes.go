package main

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/7201-12/ELS/dao"
	"github.com/7201-12/ELS/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog/log"
)

type Handler struct {
	DB *dao.Els
}

func NewHandler(pool *pgxpool.Pool) *Handler {
	dao := &dao.Els{DB: pool}
	return &Handler{DB: dao}
}

func (h *Handler) routes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:20180"},
		AllowedMethods: []string{"PUT", "POST", "DELETE", "GET", "OPTIONS"},
	}))

	// fullness + integrity = fulltegrity
	r.Get("/fulltegrity", h.GetFulltegrity)
	r.Get("/problems", h.GetProblems)
	r.Post("/calculate", h.CalculateScore)

	return r
}

func (h *Handler) GetFulltegrity(w http.ResponseWriter, r *http.Request) {
	questions := make([]*models.Question, 0)
	times := []float32{0.5, 1, 1.5}
	for _, v := range times {
		q, err := h.DB.GetQuestions(models.Fulltegrity, v)
		if err != nil {
			log.Error().Err(err).Msg("dao")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		rand.Seed(time.Now().Unix())
		questions = append(questions, q[rand.Intn(len(q))])
	}
	err := json.NewEncoder(w).Encode(questions)
	if err != nil {
		log.Error().Err(err).Msg("json encode")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetProblems(w http.ResponseWriter, r *http.Request) {
	questions := make([]*models.Question, 0)
	times := []float32{2, 3}
	for _, v := range times {
		q, err := h.DB.GetQuestions(models.Problems, v)
		if err != nil {
			log.Error().Err(err).Msg("dao")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		rand.Seed(time.Now().Unix())
		questions = append(questions, q[rand.Intn(len(q))])
	}
	err := json.NewEncoder(w).Encode(questions)
	if err != nil {
		log.Error().Err(err).Msg("json encode")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

type Answer struct {
	QuestionID string `json:"questionId"`
	AnswerID   string `json:"answerId"`
}

func (h *Handler) CalculateScore(w http.ResponseWriter, r *http.Request) {
	score := 0
	answers := make([]*Answer, 0)
	var b bytes.Buffer
	_, err := b.ReadFrom(r.Body)
	if err != nil {
		log.Error().Err(err).Msg("read from body")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(b.Bytes(), &answers)
	if err != nil {
		log.Error().Err(err).Msg("json decode")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	for _, a := range answers {
		question, err := h.DB.GetQuestion(uuid.MustParse(a.QuestionID))
		if err != nil {
			log.Error().Err(err).Msg("dao")
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		if a.AnswerID == question.Answer.ID.String() {
			score += int(question.Time * 3)
		}
	}
	err = json.NewEncoder(w).Encode(score)
	if err != nil {
		log.Error().Err(err).Msg("json encode")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
