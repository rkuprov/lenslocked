package migrations

var UserTable = `
CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	email VARCHAR(255) unique NOT NULL,
	password_hash VARCHAR(255) NOT NULL,
	created_at TIMESTAMP,
	updated_at TIMESTAMP)`
