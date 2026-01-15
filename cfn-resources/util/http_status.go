package util

import "net/http"

func StatusNotFound(resp *http.Response) bool {
	return resp != nil && resp.StatusCode == http.StatusNotFound
}

func StatusConflict(resp *http.Response) bool {
	return resp != nil && resp.StatusCode == http.StatusConflict
}

func StatusBadRequest(resp *http.Response) bool {
	return resp != nil && resp.StatusCode == http.StatusBadRequest
}

func StatusServiceUnavailable(resp *http.Response) bool {
	return resp != nil && resp.StatusCode == http.StatusServiceUnavailable
}
