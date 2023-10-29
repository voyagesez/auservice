package constants

import "github.com/go-chi/jwtauth/v5"

const (
	Google   = "google"
	Facebook = "facebook"
	Github   = "github"
	Twitter  = "twitter"
)

var (
	JWTAuthenticator *jwtauth.JWTAuth
)
