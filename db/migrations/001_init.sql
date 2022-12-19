-- +goose Up
-- +goose StatementBegin
CREATE TABLE content
(
    id      bigint      primary key,
    body    text        not null
);

COMMENT ON TABLE content IS 'the content table stores content';
-- +goose StatementEnd
