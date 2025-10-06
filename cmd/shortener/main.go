package main

import (
	"net/http"

	"github.com/vsevolod-ryzhov/urlshortner.git/internal/handler"
)

func main() {
	err := http.ListenAndServe(`:8080`, handler.MakeHandler())
	if err != nil {
		panic(err)
	}
}
