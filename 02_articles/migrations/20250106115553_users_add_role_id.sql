-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
  ADD COLUMN role_id INTEGER REFERENCES roles(id) ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
  DROP COLUMN role_id;
-- +goose StatementEnd