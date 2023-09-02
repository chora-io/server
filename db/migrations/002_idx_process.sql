-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS idx_process
(
    chain_id                varchar         not null,
    process_name            varchar         not null,
    last_block              bigint          not null,
    primary key (chain_id, process_name)
);

COMMENT ON TABLE idx_process IS 'the idx process table stores information about a process';
-- +goose StatementEnd
