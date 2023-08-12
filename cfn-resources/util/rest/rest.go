package rest

import (
	"net/http"
)

func IsHTTPSuccess(response *http.Response) bool {
	return response.StatusCode >= 200 && response.StatusCode < 300
}
