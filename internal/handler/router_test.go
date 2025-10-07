package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandlerPOSTSuccess(t *testing.T) {
	requestBody := "https://ya.ru"
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(requestBody))

	recorder := httptest.NewRecorder()

	requestsHandler(recorder, req)

	if recorder.Code != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", recorder.Code, http.StatusCreated)
	}

	contentType := recorder.Header().Get("Content-Type")
	if contentType != "text/plain" {
		t.Errorf("handler returned wrong content type: got %v want %v", contentType, "text/plain")
	}

	responseBody := recorder.Body.String()
	if !strings.HasPrefix(responseBody, "http") {
		t.Errorf("Expected to have response url started with http: got %v", responseBody)
	}
}

func TestHandlerGETNotFound(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/123", nil)

	recorder := httptest.NewRecorder()

	requestsHandler(recorder, req)

	if recorder.Code != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", recorder.Code, http.StatusNotFound)
	}
}
