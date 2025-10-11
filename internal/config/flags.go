package config

import (
	"flag"
)

var Options struct {
	AppAddress       string
	ShortenedBaseURL string
}

func ParseFlags() {
	flag.StringVar(&Options.AppAddress, "a", "localhost:8888", "The address to bind the app to")
	flag.StringVar(&Options.ShortenedBaseURL, "b", "localhost:8000", "The base url of shortened")
	flag.Parse()
}
