-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS articles (
 id SERIAL PRIMARY KEY,
 author VARCHAR(255) NOT NULL,
 title VARCHAR(255) NOT NULL,
 body TEXT NOT NULL,
 url VARCHAR(255) NOT NULL,
 provider_sign VARCHAR(255),
 published_at TIMESTAMP NOT NULL,
 created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS categories;
-- +goose StatementEnd