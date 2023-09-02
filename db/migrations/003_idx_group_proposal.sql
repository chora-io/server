-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS idx_group_proposal
(
    chain_id                varchar         not null,
    proposal_id             bigint          not null,
    proposal                jsonb           not null,
    primary key (chain_id, proposal_id)
);

COMMENT ON TABLE idx_group_proposal IS 'the final state of a group proposal for a given chain';
-- +goose StatementEnd
