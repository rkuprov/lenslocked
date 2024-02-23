package services

var (
	UserAddSQL = "INSERT INTO users (email, password_hash, created_at, updated_at) VALUES ($1, $2, $3, $3) returning id"
)
