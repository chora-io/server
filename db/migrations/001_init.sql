-- +goose Up
-- +goose StatementBegin
CREATE TABLE data
(
    iri         varchar         primary key,
    context     varchar(64)     not null,
    jsonld      varchar(1024)   unique not null
);

COMMENT ON TABLE data IS 'the data table stores linked data';
-- +goose StatementEnd
