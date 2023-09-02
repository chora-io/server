-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS idx_group_vote
(
    chain_id                varchar         not null,
    proposal_id             bigint          not null,
    voter                   text            not null,
    vote                    jsonb           not null,
    primary key (chain_id, proposal_id, voter)
);

COMMENT ON TABLE idx_group_vote IS 'the final state of a group vote for a given chain';
-- +goose StatementEnd
