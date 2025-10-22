package main

import (
	"net/http"
	"time"

	"github.com/vsevolod-ryzhov/urlshortner.git/internal/config"
	"github.com/vsevolod-ryzhov/urlshortner.git/internal/handler"
)

func main() {
	config.ParseFlags()

	srv := &http.Server{
		Addr:         config.Options.AppPort,
		Handler:      handler.MakeHandler(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
