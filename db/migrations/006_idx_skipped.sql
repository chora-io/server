-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS idx_skipped_block
(
    chain_id                varchar         not null,
    process_name            varchar         not null,
    skipped_block           bigint          not null,
    reason                  text            not null,
    primary key (chain_id, process_name, skipped_block)
);

COMMENT ON TABLE idx_skipped_block IS 'the idx skipped table stores skipped blocks to retry later';
-- +goose StatementEnd
