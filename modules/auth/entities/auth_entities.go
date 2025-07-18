package entities

type TokenRequest struct {
	GrantType    string `db:"grant_type" form:"grant_type" json:"grant_type" validate:"required"`
	ClientID     string `db:"client_id" form:"client_id" json:"client_id" validate:"required"`
	ClientSecret string `db:"client_secret" form:"client_secret" json:"client_secret" validate:"required"`
}

type ClientData struct {
	Id           int    `db:"id" json:"id"`
	SystemSource string `db:"system_source" json:"system_source"`
	IsActive     bool   `db:"is_active" json:"is_active"`
	TokenRequest
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   uint32 `json:"expires_in"`
	TokenType   string `json:"token_type"`
}
