package model

type JWTToken struct {
	Token       string `json:"token"`
	Description string `json:"description"`
}
