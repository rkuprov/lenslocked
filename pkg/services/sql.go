package services

var (
	UserAddSQL             = "INSERT INTO users (email, password_hash, created_at, updated_at) VALUES ($1, $2, $3, $3) returning id"
	UserGetPasswordHashSQL = "SELECT password_hash FROM users WHERE email = $1"
	SessionAddSQL          = "INSERT INTO sessions (user_id, token_hash, created_at) VALUES ($1, $2, $3)"
)
