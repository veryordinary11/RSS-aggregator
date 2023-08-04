package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey returns the API key from the request headers
func GetAPIKey(headers http.Header) (string, error) {
	apiKey := headers.Get("Authorization")
	if apiKey == "" {
		return "", errors.New("missing API key")
	}

	vals := strings.Split(apiKey, " ")
	if len(vals) != 2 {
		return "", errors.New("invalid API key")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("invalid API key")
	}

	return vals[1], nil
}
