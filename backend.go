package ssosdk

import (
	"fmt"
	"net/http"
)

// ---------------------------------------------------------------------------
// IAM / BO
// ---------------------------------------------------------------------------

// ExchangeCode exchanges an authorization code for JWT claims.
//
// GET /v1/backend/iam/bo/exchange-code/{code}
// No bearer token required (the code itself is the credential).
func (c *Client) ExchangeCode(code string) (ExchangeCodeResponse, error) {
	path := fmt.Sprintf("/v1/backend/iam/bo/exchange-code/%s", code)
	var resp ExchangeCodeResponse
	if err := c.do(http.MethodGet, path, nil, &resp, false); err != nil {
		return nil, err
	}
	return resp, nil
}

// RefreshToken exchanges the current bearer token for a new access token.
//
// GET /v1/backend/iam/bo/refresh-token
// Requires bearer token.
func (c *Client) RefreshToken() (*RefreshTokenResponse, error) {
	var resp RefreshTokenResponse
	if err := c.do(http.MethodGet, "/v1/backend/iam/bo/refresh-token", nil, &resp, true); err != nil {
		return nil, err
	}
	return &resp, nil
}

// VerifyToken verifies the current bearer token and returns its claims.
//
// GET /v1/backend/iam/bo/verify-token
// Requires bearer token.
func (c *Client) VerifyToken() (VerifyTokenResponse, error) {
	var resp VerifyTokenResponse
	if err := c.do(http.MethodGet, "/v1/backend/iam/bo/verify-token", nil, &resp, true); err != nil {
		return nil, err
	}
	return resp, nil
}

// ---------------------------------------------------------------------------
// User management
// ---------------------------------------------------------------------------

// CreateUser creates a new user in the authenticated pool.
//
// POST /v1/backend/user
// Requires bearer token (pool context is derived from the token).
func (c *Client) CreateUser(req CreateUserRequest) (*CreateUserResponse, error) {
	var resp CreateUserResponse
	if err := c.do(http.MethodPost, "/v1/backend/user", req, &resp, true); err != nil {
		return nil, err
	}
	return &resp, nil
}
