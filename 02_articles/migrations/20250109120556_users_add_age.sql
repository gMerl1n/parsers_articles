-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
  ADD COLUMN age INTEGER;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
  DROP COLUMN age;
-- +goose StatementEnd