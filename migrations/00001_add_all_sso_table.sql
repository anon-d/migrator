-- +goose Up
-- +goose StatementBegin
SELECT
    'up SQL query';

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    email VARCHAR(320) NOT NULL UNIQUE,
    passhash BYTEA NOT NULL UNIQUE
);

-- CREATE UNIQUE INDEX idx_users_email_unique ON users (uuid);
CREATE TABLE IF NOT EXISTS apps (
    id SERIAL PRIMARY KEY,
    name VARCHAR(128) NOT NULL UNIQUE
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
SELECT
    'down SQL query';

DROP TABLE IF EXISTS users;

-- DROP INDEX IF EXISTS idx_users_email_unique;
DROP TABLE IF EXISTS apps;

-- +goose StatementEnd
