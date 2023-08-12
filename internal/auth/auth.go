package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts an API key from the headers of and http request
// Example:
// Authorization: ApiKey *Insert apikey here*
func GetAPIKey(headers http.Header) (string, error) {
	authkey := headers.Get("Authorization")
	if authkey == "" {
		return "", errors.New("no authentication Info")
	}

	vals := strings.Split(authkey, " ")

	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of auth header")
	}

	return vals[1], nil
}
