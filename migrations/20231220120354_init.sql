-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS auth.user(
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    deleted_at TIMESTAMP,
    user_id SERIAL PRIMARY KEY,
    username VARCHAR NOT NULL,
    full_name VARCHAR,
    password VARCHAR NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE auth.user;
-- +goose StatementEnd
