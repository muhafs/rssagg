package auth

import (
	"errors"
	"net/http"
	"strings"
)

/*
* GetApiKey exctract an API Key from
* the headers of an HTTP request
*
* - Example:
* 	- Authorization: ApiKey {insert apiKey here}
* 	- "ApiKey d38ewe452cw2b5733w553vbeef5cf23ea6344b2823c75f56a2ed25697f869889"
 */
func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authontication info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of auth header")
	}

	return vals[1], nil
}
