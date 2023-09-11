-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS auth_user
(
    address                 varchar         not null,
    created_at              timestamp       not null,
    last_authenticated      timestamp       not null,
    primary key (address)
);

COMMENT ON TABLE auth_user IS 'authenticated user';
-- +goose StatementEnd
