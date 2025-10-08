package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vsevolod-ryzhov/urlshortner.git/internal/handler"
	"github.com/vsevolod-ryzhov/urlshortner.git/internal/service"
)

func testRequest(t *testing.T, ts *httptest.Server, method string, path string, body io.Reader) (*http.Response, string) {
	req, err := http.NewRequest(method, ts.URL+path, body)
	require.NoError(t, err)

	resp, err := ts.Client().Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	return resp, string(respBody)
}

func TestRouter(t *testing.T) {
	ts := httptest.NewServer(handler.MakeHandler())
	defer ts.Close()
	originalURL := "https://ya.ru"

	resp, get := testRequest(t, ts, "POST", "", strings.NewReader(originalURL))
	defer resp.Body.Close()
	code := strings.Replace(get, ts.URL+"/", "", -1)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
	assert.Equal(t, service.CreateShortURL(originalURL), code)

	shortenedCode := strings.Replace(get, ts.URL, "", -1)
	getResp, _ := testRequest(t, ts, "GET", shortenedCode, nil)
	getResp.Body.Close()
	assert.Equal(t, http.StatusOK, getResp.StatusCode) // status code from destination URL on which script is redirected
}
