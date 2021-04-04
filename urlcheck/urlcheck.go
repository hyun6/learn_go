package urlcheck

import (
	"errors"
	"net/http"
)

type Result struct {
	Url string
	Status string
}


var errRequestFailed = errors.New("request failed")

// chan<- : send only channel
func Check(url string, c chan<- Result) {
	resp, err := http.Get(url);
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}

	c <- Result{url, status}
}