package auth

import (
	"errors"
	"net/http"
	"strings"
)

// extract API key from headers
// exampel -- Authorization: Apikey {API_KEY_HERE}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")

	if val == "" {
		return "", errors.New("no auth info found")
	}

	vals := strings.Split(val, " ")

	if len(vals) != 2 {
		return "", errors.New("malformed Header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("no auth info found")
	}

	return vals[1], nil

}
