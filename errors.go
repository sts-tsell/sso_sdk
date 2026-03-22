package ssosdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// ErrNoToken is returned when an authenticated call is made without a bearer token.
var ErrNoToken = errors.New("ssosdk: bearer token is not set")

// APIError represents an error response from the SSO API.
type APIError struct {
	StatusCode int
	Body       string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("ssosdk: API error %d: %s", e.StatusCode, e.Body)
}

func parseErrorResponse(resp *http.Response) error {
	body, _ := io.ReadAll(resp.Body)

	// Try to extract a human-readable message from JSON.
	var payload struct {
		Error string `json:"error"`
	}
	if json.Unmarshal(body, &payload) == nil && payload.Error != "" {
		return &APIError{StatusCode: resp.StatusCode, Body: payload.Error}
	}

	return &APIError{StatusCode: resp.StatusCode, Body: string(body)}
}
