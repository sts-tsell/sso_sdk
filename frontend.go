package ssosdk

import (
	"fmt"
	"net/http"
	"net/url"
)

// ---------------------------------------------------------------------------
// SMS Auth
// ---------------------------------------------------------------------------

// CreateSMSAuth initiates an SMS OTP authentication request.
//
// POST /v1/frontend/sms-auth
// No bearer token required.
func (c *Client) CreateSMSAuth(req CreateSMSAuthRequest) (*CreateSMSAuthResponse, error) {
	var resp CreateSMSAuthResponse
	if err := c.do(http.MethodPost, "/v1/frontend/sms-auth", req, &resp, false); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ConfirmSMSAuth confirms an SMS OTP code and returns an access token.
//
// PATCH /v1/frontend/sms-auth
// No bearer token required.
func (c *Client) ConfirmSMSAuth(req ConfirmSMSAuthRequest) (*ConfirmSMSAuthResponse, error) {
	var resp ConfirmSMSAuthResponse
	if err := c.do(http.MethodPatch, "/v1/frontend/sms-auth", req, &resp, false); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ---------------------------------------------------------------------------
// Me
// ---------------------------------------------------------------------------

// GetMe returns the authenticated user's profile for the given pool.
//
// GET /v1/frontend/me?pool_id={poolID}
// Requires bearer token.
func (c *Client) GetMe(poolID string) (*GetMeResponse, error) {
	path := fmt.Sprintf("/v1/frontend/me?%s", url.Values{"pool_id": {poolID}}.Encode())
	var resp GetMeResponse
	if err := c.do(http.MethodGet, path, nil, &resp, true); err != nil {
		return nil, err
	}
	return &resp, nil
}

// VerifyAPIKey verifies an API key for the authenticated user and pool.
//
// GET /v1/frontend/me/verify-api-key?pool_id={poolID}&api_key={apiKey}
// Requires bearer token.
func (c *Client) VerifyAPIKey(poolID, apiKey string) (*VerifyAPIKeyResponse, error) {
	q := url.Values{"pool_id": {poolID}, "api_key": {apiKey}}
	path := fmt.Sprintf("/v1/frontend/me/verify-api-key?%s", q.Encode())
	var resp VerifyAPIKeyResponse
	if err := c.do(http.MethodGet, path, nil, &resp, true); err != nil {
		return nil, err
	}
	return &resp, nil
}
