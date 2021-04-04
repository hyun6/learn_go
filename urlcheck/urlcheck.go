package urlcheck

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("request failed")

func Check(url string) error {
	resp, err := http.Get(url);
	if err != nil || resp.StatusCode >= 400 {
		fmt.Printf("%s request error: status=%d\n", url, resp.StatusCode)
		return errRequestFailed
	}
	return nil;
}