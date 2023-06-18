-- name: GetData :one
select iri, jsonld from data where iri=$1;

-- name: SelectIdxGroupProposal :one
select proposal from idx_group_proposal where chain_id=$1 and proposal_id=$2;

-- name: SelectIdxGroupProposals :many
select proposal from idx_group_proposal where chain_id=$1;

-- name: SelectIdxProcessLastBlock :one
select last_block from idx_process where chain_id=$1 and process_name=$2;
