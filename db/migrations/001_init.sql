-- +goose Up
-- +goose StatementBegin
CREATE TABLE data
(
    iri         varchar         primary key,
    jsonld      jsonb           unique not null
);

COMMENT ON TABLE data IS 'the data table stores linked data';
-- +goose StatementEnd
