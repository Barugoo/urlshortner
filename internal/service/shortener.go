package service

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"strings"
)

var urlStorage = make(map[string]string)

func CreateShortUrl(url string) string {
	shortId := generateShortId(url)

	urlStorage[shortId] = url

	return shortId
}

func Get(id string) (string, error) {
	url, exists := urlStorage[id]
	if !exists {
		return "", errors.New("url not found")
	}

	return url, nil
}

func generateShortId(originalUrl string) string {

	for id, url := range urlStorage {
		if url == originalUrl {
			return id
		}
	}

	hash := sha256.Sum256([]byte(originalUrl))
	shortID := base64.URLEncoding.EncodeToString(hash[:8])
	return strings.TrimRight(shortID, "=")
}
