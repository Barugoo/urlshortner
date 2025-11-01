package main

import (
	"net/http"
	"time"

	"github.com/vsevolod-ryzhov/urlshortner.git/internal/config"
	"github.com/vsevolod-ryzhov/urlshortner.git/internal/handler"
	"github.com/vsevolod-ryzhov/urlshortner.git/internal/logger"
)

func main() {
	config.ParseFlags()

	if err := logger.Initialize(config.Options.FlagLogLevel); err != nil {
		panic(err)
	}

	srv := &http.Server{
		Addr:         config.Options.AppPort,
		Handler:      logger.WithLogging(handler.GzipMiddleware(handler.MakeHandler())),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
