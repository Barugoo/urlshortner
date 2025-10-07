package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/vsevolod-ryzhov/urlshortner.git/internal/service"
)

func requestsHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		body := make([]byte, req.ContentLength)

		_, err := req.Body.Read(body)
		if err != nil && err.Error() != "EOF" {
			http.Error(res, "Bad request", http.StatusBadRequest)
			return
		}

		url := string(body)

		shortened := service.CreateShortURL(url)
		fullURL := fmt.Sprintf("http://%s/%s", req.Host, shortened)
		res.WriteHeader(http.StatusCreated)
		res.Write([]byte(fullURL))
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

		res.Header().Add("Location", url)
		res.WriteHeader(http.StatusTemporaryRedirect)
	default:
		http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func MakeHandler() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc(`/`, requestsHandler)

	return mux
}
