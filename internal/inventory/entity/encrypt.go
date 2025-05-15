package entity

type CodeMessage struct {
	CodeChallenge       string `json:"codeChallenge"`
	CodeChallengeMethod string `json:"codeChallengeMethod"`
	AuthorizationCode   string `json:"authorizationCode"`
}
