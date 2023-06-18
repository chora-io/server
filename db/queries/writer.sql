-- name: InsertData :exec
insert into data (iri, jsonld) values ($1, $2);

-- name: InsertIdxGroupProposal :exec
insert into idx_group_proposal (chain_id, proposal_id, proposal) values ($1, $2, $3);

-- name: InsertIdxProcessLastBlock :exec
insert into idx_process (chain_id, process_name, last_block) values ($1, $2, $3);

-- name: UpdateIdxProcessLastBlock :exec
update idx_process set last_block=$3 where chain_id=$1 and process_name=$2;
