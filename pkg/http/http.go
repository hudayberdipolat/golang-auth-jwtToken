package http

import (
	"net/http"
	"time"
)

func NewHttpClient() *http.Client {
	client := &http.Client{
		Timeout: 2 * time.Minute,
	}
	return client
}
