package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Exmple:
// Authorizetion: ApiKey  {insert apiKey here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("authorization header is missing")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 || vals[0] != "ApiKey" {
		return "", errors.New("authorization header is malformed")
	}

	return vals[1], nil
}
