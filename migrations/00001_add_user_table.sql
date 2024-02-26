-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	email VARCHAR(255) unique NOT NULL,
	password_hash VARCHAR(255) NOT NULL,
	created_at TIMESTAMP,
	updated_at TIMESTAMP);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
