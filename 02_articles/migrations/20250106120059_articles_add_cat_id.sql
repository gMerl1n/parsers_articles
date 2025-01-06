-- +goose Up
-- +goose StatementBegin
ALTER TABLE articles
  ADD COLUMN category_id INTEGER REFERENCES categories(id) ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE articles
  DROP COLUMN category_id;
-- +goose StatementEnd