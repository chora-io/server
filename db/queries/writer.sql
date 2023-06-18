-- name: PostData :exec
insert into data (iri, jsonld) values ($1, $2);

-- name: AddIdxGroupProposal :exec
insert into idx_group_proposal (chain_id, proposal_id, proposal) values ($1, $2, $3);

-- name: UpdateIdxProcessLastBlock :exec
update idx_process set last_block=$3 where chain_id=$1 and process_name=$2;
