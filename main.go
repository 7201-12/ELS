package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	for i := 0; i < 8; i++ {
		fmt.Println(uuid.New())
	}
	connStr := "postgres://postgres:postgres@localhost:5430/els?sslmode=disable"
	pool, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal().Err(err).Msg("open migrations db connection")
	}
	h := NewHandler(pool)
	r := h.routes()
	fmt.Println("Listen on :4000..........")
	err = http.ListenAndServe(":4000", r)
	if err != nil {
		log.Fatal().Err(err).Msg("listen and serve error")
	}
}
