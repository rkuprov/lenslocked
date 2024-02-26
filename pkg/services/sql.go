package services

var (
	UserAddSQL             = "INSERT INTO users (email, password_hash, created_at, updated_at) VALUES ($1, $2, $3, $3) returning id"
	UserGetPasswordHashSQL = "SELECT id, password_hash FROM users WHERE email = $1"
	SessionAddSQL          = `INSERT INTO sessions (user_id, token_hash, created_at) VALUES ($1, $2, $3) 
								ON CONFLICT (user_id) DO UPDATE 
								SET token_hash=$2, created_at = $3 
								returning id`
	SessionGetUserSQL = "SELECT u.id, u.email FROM users u JOIN sessions s ON u.id = s.user_id WHERE s.token_hash = $1"
	SessionDeleteSQL  = "DELETE FROM sessions WHERE token_hash = $1"
)
