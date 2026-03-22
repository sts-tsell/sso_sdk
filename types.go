package ssosdk

// ---------------------------------------------------------------------------
// Frontend — SMS auth
// ---------------------------------------------------------------------------

// CreateSMSAuthRequest is the payload for POST /v1/frontend/sms-auth.
type CreateSMSAuthRequest struct {
	Username string `json:"username"`
	PoolID   string `json:"pool_id"`
	// Silent — when true the OTP is delivered silently (e.g. via Telegram).
	Silent *bool `json:"silent,omitempty"`
}

// CreateSMSAuthResponse is returned by POST /v1/frontend/sms-auth.
type CreateSMSAuthResponse struct {
	ID string `json:"id"`
	// Code is only populated when the server is configured to return it directly.
	Code *int `json:"code,omitempty"`
}

// ConfirmSMSAuthRequest is the payload for PATCH /v1/frontend/sms-auth.
type ConfirmSMSAuthRequest struct {
	AuthRequestID string `json:"auth_request_id"`
	Code          int    `json:"code"`
}

// ConfirmSMSAuthResponse is returned by PATCH /v1/frontend/sms-auth.
type ConfirmSMSAuthResponse struct {
	Token  string `json:"token"`
	UserID string `json:"user_id"`
	PoolID string `json:"pool_id"`
}

// ---------------------------------------------------------------------------
// Frontend — Me
// ---------------------------------------------------------------------------

// GetMeResponse is returned by GET /v1/frontend/me.
type GetMeResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Role     int    `json:"role"`
	Token    string `json:"token"`
}

// VerifyAPIKeyResponse is returned by GET /v1/frontend/me/verify-api-key.
type VerifyAPIKeyResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Role     int    `json:"role"`
	Token    string `json:"token"`
}

// ---------------------------------------------------------------------------
// Backend — IAM / BO
// ---------------------------------------------------------------------------

// ExchangeCodeResponse is returned by GET /v1/backend/iam/bo/exchange-code/{code}.
// The server returns the raw JWT claims object, so this is a generic map.
type ExchangeCodeResponse map[string]interface{}

// RefreshTokenResponse is returned by GET /v1/backend/iam/bo/refresh-token.
type RefreshTokenResponse struct {
	AccessToken string `json:"access_token"`
}

// VerifyTokenResponse is returned by GET /v1/backend/iam/bo/verify-token.
// The server returns the raw JWT claims object.
type VerifyTokenResponse map[string]interface{}

// ---------------------------------------------------------------------------
// Backend — User
// ---------------------------------------------------------------------------

// CreateUserRequest is the payload for POST /v1/backend/user.
type CreateUserRequest struct {
	Username string `json:"username"`
}

// CreateUserResponse is returned by POST /v1/backend/user.
type CreateUserResponse struct {
	Token  string `json:"token"`
	UserID string `json:"user_id"`
	PoolID string `json:"pool_id"`
}
