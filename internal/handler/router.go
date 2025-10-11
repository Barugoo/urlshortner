package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/vsevolod-ryzhov/urlshortner.git/internal/config"
	"github.com/vsevolod-ryzhov/urlshortner.git/internal/service"
)

func handleCreateLink(res http.ResponseWriter, req *http.Request) {
	body := make([]byte, req.ContentLength)

	_, err := req.Body.Read(body)
	if err != nil && err.Error() != "EOF" {
		http.Error(res, "Bad request", http.StatusBadRequest)
		return
	}

	url := string(body)

	shortened := service.CreateShortURL(url)
	shortenedURL := fmt.Sprintf("%s/%s", config.Options.ShortenedBaseURL, shortened)
	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(shortenedURL))
}

func handleGetLink(res http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/")

	if id == "" {
		http.Error(res, "ID is not specified", http.StatusBadRequest)
		return
	}

	url, err := service.GetURL(id)
	if err != nil {
		http.Error(res, "URL not found", http.StatusNotFound)
		return
	}
	res.Header().Add("Location", url)
	res.WriteHeader(http.StatusTemporaryRedirect)
}

func MakeHandler() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/{link}", handleGetLink)
	r.Post("/", handleCreateLink)

	return r
}
