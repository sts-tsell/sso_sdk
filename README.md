# sso_sdk

Go SDK for the [single_login_api](https://github.com/sts-tsell/single_login_api) SSO service.

## Installation

```bash
go get github.com/sts-tsell/sso_sdk
```

## Quick start

```go
package main

import (
    "fmt"
    ssosdk "github.com/sts-tsell/sso_sdk"
)

func main() {
    // Create a client — base URL is required, bearer token is optional at construction time.
    client := ssosdk.New("https://sso.example.com")

    // 1. Initiate SMS OTP login
    authResp, err := client.CreateSMSAuth(ssosdk.CreateSMSAuthRequest{
        Username: "john@example.com",
        PoolID:   "550e8400-e29b-41d4-a716-446655440000",
    })
    if err != nil {
        panic(err)
    }

    // 2. Confirm the OTP code the user received
    tokenResp, err := client.ConfirmSMSAuth(ssosdk.ConfirmSMSAuthRequest{
        AuthRequestID: authResp.ID,
        Code:          123456,
    })
    if err != nil {
        panic(err)
    }

    fmt.Println("access token:", tokenResp.Token)

    // 3. Use the token for authenticated calls
    client.SetBearerToken(tokenResp.Token)

    me, err := client.GetMe(tokenResp.PoolID)
    if err != nil {
        panic(err)
    }
    fmt.Printf("logged in as %s (role %d)\n", me.Username, me.Role)
}
```

## Configuration

| Option | Description |
|--------|-------------|
| `ssosdk.New(baseURL)` | Required. Base URL of the SSO API, no trailing slash. |
| `ssosdk.WithBearerToken(token)` | Set bearer token at construction time. |
| `ssosdk.WithHTTPClient(hc)` | Replace the default `http.Client` (30 s timeout). |
| `client.SetBearerToken(token)` | Update the bearer token at runtime (e.g. after login). |

## API coverage

### Frontend

| Method | Endpoint | Auth |
|--------|----------|------|
| `CreateSMSAuth` | `POST /v1/frontend/sms-auth` | — |
| `ConfirmSMSAuth` | `PATCH /v1/frontend/sms-auth` | — |
| `GetMe` | `GET /v1/frontend/me` | Bearer |
| `VerifyAPIKey` | `GET /v1/frontend/me/verify-api-key` | Bearer |

### Backend

| Method | Endpoint | Auth |
|--------|----------|------|
| `ExchangeCode` | `GET /v1/backend/iam/bo/exchange-code/{code}` | — |
| `RefreshToken` | `GET /v1/backend/iam/bo/refresh-token` | Bearer |
| `VerifyToken` | `GET /v1/backend/iam/bo/verify-token` | Bearer |
| `CreateUser` | `POST /v1/backend/user` | Bearer |

## Error handling

All methods return an `error`. API-level errors are returned as `*ssosdk.APIError`:

```go
var apiErr *ssosdk.APIError
if errors.As(err, &apiErr) {
    fmt.Println(apiErr.StatusCode, apiErr.Body)
}
```

`ssosdk.ErrNoToken` is returned when an authenticated method is called without a bearer token set.
