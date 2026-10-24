package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts an API Key from
// the headers of an HTTP request
// Example:
// Authorization: ApiKey {apiKey here}
func GetAPIKey(header http.Header) (string, error) {
	val := header.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication info found")
	}
	v := strings.Split(val, " ")
	// log.Println(v)
	if len(v) < 2 {
		return "", errors.New("malformed auth header found")
	}
	if v[0] != "ApiKey" {
		return "", errors.New("malformed first part of auth header")
	}
	return v[1], nil
}
