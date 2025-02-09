-- +goose Up
ALTER TABLE users ADD COLUMN counter INTEGER NOT NULL DEFAULT 0;

-- +goose Down
ALTER TABLE users DROP COLUMN counter;
