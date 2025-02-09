-- +goose Up
ALTER TABLE feeds ADD CONSTRAINT feeds_url_key UNIQUE (url);

-- +goose Down
ALTER TABLE feeds DROP CONSTRAINT feeds_url_key;