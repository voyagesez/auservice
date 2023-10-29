package utils

import (
	"context"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/voyagesez/auservice/src/constants"
	"github.com/voyagesez/auservice/src/internals/dtos"
)

type JWTAccessTokenDecode struct {
	UID   pgtype.UUID `json:"uid"`
	Email string      `json:"email"`
}

func GenerateAccessToken(payload dtos.UserAccountResponse) (*dtos.TokenResponse, error) {
	expiredAt := time.Now().Add(time.Minute * 60)
	claims := map[string]interface{}{
		"uid":   payload.UID,
		"email": payload.Email,
	}

	jwtauth.SetExpiry(claims, expiredAt)
	jwtauth.SetIssuedAt(claims, time.Now())
	_, tokenString, err := constants.JWTAuthenticator.Encode(claims)
	if err != nil {
		return nil, err
	}

	return &dtos.TokenResponse{
		AccessToken: tokenString,
		TokenType:   "Bearer",
		ExpiresIn:   expiredAt.Unix(),
	}, nil
}

func DecodeAccessToken(tokenString string) (*JWTAccessTokenDecode, error) {
	jwtToken, err := constants.JWTAuthenticator.Decode(tokenString)
	if err != nil {
		return nil, err
	}
	claims, err := jwtToken.AsMap(context.Background())
	if err != nil {
		return nil, err
	}
	return mapTokeTokenClaimsStruct(claims)
}

func mapTokeTokenClaimsStruct(claims map[string]interface{}) (*JWTAccessTokenDecode, error) {
	return &JWTAccessTokenDecode{}, nil
}
