package dtos

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	IdToken      string `json:"id_token"`
	ExpiresIn    int64  `json:"expires_in"`
}
