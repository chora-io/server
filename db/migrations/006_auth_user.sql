-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS auth_user
(
    id                      varchar         primary key,
    address                 varchar         unique,
    email                   varchar         unique,
    username                varchar         unique,
    created_at              timestamp       not null
);

COMMENT ON TABLE auth_user IS 'authenticated user';
-- +goose StatementEnd
