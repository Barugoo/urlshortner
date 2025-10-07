package service

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"strings"
)

var urlStorage = make(map[string]string)

func CreateShortURL(url string) string {
	shortID := generateShortID(url)

	urlStorage[shortID] = url

	return shortID
}

func GetURL(id string) (string, error) {
	url, exists := urlStorage[id]
	if !exists {
		return "", errors.New("url not found")
	}

	return url, nil
}

func generateShortID(originalURL string) string {

	for id, url := range urlStorage {
		if url == originalURL {
			return id
		}
	}

	hash := sha256.Sum256([]byte(originalURL))
	shortID := base64.URLEncoding.EncodeToString(hash[:8])
	return strings.TrimRight(shortID, "=")
}
