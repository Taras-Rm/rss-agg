package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey returns ApiKey from HTTP
// request headers
// Example:
// Authorization: ApiKey {api key here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authorization info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("invalid authorization header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("invalid first part of authorization header")
	}

	return vals[1], nil
}
