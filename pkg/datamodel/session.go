package datamodel

type Session struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	TokenHash string `json:"token_hash"`
}
