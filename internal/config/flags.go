package config

import (
	"flag"
)

var Options struct {
	AppPort          string
	ShortenedBaseURL string
}

func ParseFlags() {
	flag.StringVar(&Options.AppPort, "a", "localhost:8080", "The address to bind the app to")
	flag.StringVar(&Options.ShortenedBaseURL, "b", "localhost:8080", "The base url of shortened")
	flag.Parse()
}
