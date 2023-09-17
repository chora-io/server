-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS auth_user
(
    id                      varchar         primary key,
    email                   varchar         unique,
    address                 varchar         unique,
    username                varchar         unique,
    created_at              timestamp       not null
);

COMMENT ON TABLE auth_user IS 'authenticated user';
-- +goose StatementEnd
