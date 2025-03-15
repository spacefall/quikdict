package utils

import (
	"errors"
	"io"
	"net/http"
)

func GetHttp(url string) ([]byte, error) {
	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	// Return definition not found when not 200 (even if not quite universally correct)
	if r.StatusCode != 200 {
		return nil, errors.New("definition not found")
	}

	// Parse the response body
	return io.ReadAll(r.Body)
}
