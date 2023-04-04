-- +goose Up
CREATE TABLE IF NOT EXISTS "users"
(
    id            bigserial,
    created_at    timestamp with time zone NOT NULL,
    updated_at    timestamp with time zone NOT NULL,
    deleted_at    timestamp with time zone DEFAULT NULL,
    "name"        text                     NOT NULL,
    "email"       text                     NOT NULL,
    "password"    text                     NOT NULL,
    "last_login"  timestamp with time zone DEFAULT NULL,
    CONSTRAINT UC_User unique (email, deleted_at),
    PRIMARY KEY ("id")
);

-- +goose Down
DROP TABLE users;
