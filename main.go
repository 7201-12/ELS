package main

import (
	"context"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	connStr := "postgres://postgres:postgres@localhost:5430/els?sslmode=disable"
	pool, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal().Err(err).Msg("open migrations db connection")
	}
	h := NewHandler(pool)
	r := h.routes()
	err = http.ListenAndServe(":4000", r)
	if err != nil {
		log.Fatal().Err(err).Msg("listen and serve error")
	}
}
