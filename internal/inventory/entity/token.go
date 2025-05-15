package entity

type TokenRequest struct {
	Code          string `json:"code"`
	CodeChallenge string `json:"codeChallenge"`
}

type TokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    int    `json:"expiresIn"`
	Scope        string `json:"scope"`
	TokenType    string `json:"tokenType"`
}
