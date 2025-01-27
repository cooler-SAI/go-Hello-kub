package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go-Hello-kub/handlers"
	"net/http"
	"os"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	log.Info().Msg("Loading Simple Server...")
	http.HandleFunc("/", handlers.HelloHandler)
	log.Info().Msg("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
