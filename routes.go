package main

import (
	"encoding/json"
	"net/http"

	"github.com/7201-12/ELS/dao"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog/log"
)

type Handler struct {
	DB *dao.Els
}

func NewHandler(pool *pgxpool.Pool) *Handler {
	return &Handler{}
}

func (h *Handler) routes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:20180"},
		AllowedMethods: []string{"PUT", "POST", "DELETE", "GET", "OPTIONS"},
	}))

	r.Get("/questions", h.GetQuestions)
	r.Post("/questions", h.CalculateScore)

	return r
}

func (h *Handler) GetQuestions(w http.ResponseWriter, r *http.Request) {
	questions, err := h.DB.GetQuestionsByType(1)
	if err != nil {
		log.Error().Err(err).Msg("dao")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(questions)
	if err != nil {
		log.Error().Err(err).Msg("json encode")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) CalculateScore(w http.ResponseWriter, r *http.Request) {

}
