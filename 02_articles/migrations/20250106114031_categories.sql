-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS categories (
 id SERIAL PRIMARY KEY,
 name VARCHAR(255) NOT NULL,
 url VARCHAR(255) NOT NULL,
 provider_sign VARCHAR(255),
 created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS categories;
-- +goose StatementEnd