package api

import (
	"net/http"
	"time"
)

var client = &http.Client{
	Timeout: time.Second * 10,
}

func Do(req *http.Request) (*http.Response, error) {
	return client.Do(req)
}
