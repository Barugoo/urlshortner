package main

import (
	"net/http"

	"github.com/vsevolod-ryzhov/urlshortner.git/internal/config"
	"github.com/vsevolod-ryzhov/urlshortner.git/internal/handler"
)

func main() {
	config.ParseFlags()

	err := http.ListenAndServe(config.Options.AppAddress, handler.MakeHandler())
	if err != nil {
		panic(err)
	}
}
