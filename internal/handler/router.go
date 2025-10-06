package handler

import (
	"net/http"
	"strings"

	"github.com/vsevolod-ryzhov/urlshortner.git/internal/service"
)

func requestsHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		url := req.FormValue("url")

		if url == "" {
			http.Error(res, "URL couldn't be empty", http.StatusBadRequest)
		}

		shortened := service.CreateShortUrl(url)
		res.WriteHeader(http.StatusCreated)
		res.Write([]byte(shortened))
	case http.MethodGet:
		id := strings.TrimPrefix(req.URL.Path, "/")

		if id == "" {
			http.Error(res, "ID is not specified", http.StatusBadRequest)
			return
		}

		url, err := service.Get(id)
		if err != nil {
			http.Error(res, "URL not found", http.StatusNotFound)
			return
		}

		res.WriteHeader(http.StatusTemporaryRedirect)
		res.Write([]byte(url))
	default:
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func MakeHandler() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc(`/`, requestsHandler)

	return mux
}
