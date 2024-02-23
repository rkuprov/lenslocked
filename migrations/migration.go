package migrations

var UserTable = `
CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	email VARCHAR(255),
	password_hash VARCHAR(255),
	created_at TIMESTAMP,
	updated_at TIMESTAMP)`
