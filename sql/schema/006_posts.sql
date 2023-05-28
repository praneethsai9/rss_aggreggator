-- +goose Up

CREATE TABLE posts (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title TEXT NOT NULL,
    description text,
    published_at TIMESTAMP not null,
    url  text NOT NUll unique,
    feed_id UUID not null references feeds(id) on delete cascade
);
-- +goose Down
DROP TABLE posts;