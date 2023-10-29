package dtos

import "github.com/jackc/pgx/v5/pgtype"

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
}

type CookieResponse struct {
	AccessToken  string `json:"access_token"`
	SessionId    string `json:"session_id"`
	UserCacheKey string `json:"user_cache"`
}

type UserAccountResponse struct {
	UID      pgtype.UUID `json:"uid"`
	Email    string      `json:"email"`
	FullName string      `json:"full_name"`
	Avatar   string      `json:"avatar"`
}
