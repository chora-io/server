-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS data
(
    iri         varchar         primary key,
    jsonld      jsonb           unique not null
);

COMMENT ON TABLE data IS 'the data table stores linked data';
-- +goose StatementEnd
