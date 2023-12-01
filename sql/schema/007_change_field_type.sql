-- +goose Up

ALTER TABLE posts DROP CONSTRAINT posts_url_key;

-- +goose Down

ALTER TABLE posts ADD CONSTRAINT posts_url_key UNIQUE (url);
