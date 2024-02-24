package migrations

var UserTable = `
CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	email VARCHAR(255) unique NOT NULL,
	password_hash VARCHAR(255) NOT NULL,
	created_at TIMESTAMP,
	updated_at TIMESTAMP)`

var SessionTable = `
CREATE TABLE IF NOT EXISTS sessions (
	id SERIAL PRIMARY KEY,
	user_id INT NOT NULL UNIQUE REFERENCES users(id) ON DELETE CASCADE,
	token_hash VARCHAR(255) NOT NULL UNIQUE,
	created_at TIMESTAMP)
`
