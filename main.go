package main

import (
	"errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go-Hello-kub/handlers"
	"go-Hello-kub/tools"
	"net/http"
	"os"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	log.Info().Msg("Loading Simple Server...")
	http.HandleFunc("/", handlers.HelloHandler)

	server := &http.Server{
		Addr: ":8080",
	}

	go func() {
		log.Info().Msg("Server is running on http://localhost:8080")
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Msg("Server error")
		}
	}()

	tools.StopApp(server)
}
