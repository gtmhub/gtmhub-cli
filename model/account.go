package model

type AccountResolveResponse struct {
	User UserResponse `json:"user"`
	Account AccountResponse `json:"account"`
}

type UserResponse struct {
	ID string `json:"id"`
	ClientID string `json:"clientId"`
}

type AccountResponse struct {
	ID string `json:"id"`
	Domain string `json:"domain"`
}
