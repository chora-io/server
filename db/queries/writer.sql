-- name: InsertAuthUser :exec
insert into auth_user (address, created_at, last_authenticated) values ($1, now(), now());

-- name: UpdateAuthUserLastAuthenticated :exec
update auth_user set last_authenticated=now() where address=$1;

-- name: InsertData :exec
insert into data (iri, jsonld) values ($1, $2);

-- name: InsertIdxGroupProposal :exec
insert into idx_group_proposal (chain_id, proposal_id, group_id, proposal) values ($1, $2, $3, $4);

-- name: UpdateIdxGroupProposal :exec
update idx_group_proposal set proposal=$3 where chain_id=$1 and proposal_id=$2;

-- name: InsertIdxGroupVote :exec
insert into idx_group_vote (chain_id, proposal_id, voter, vote) values ($1, $2, $3, $4);

-- name: UpdateIdxGroupVote :exec
update idx_group_vote set vote=$4 where chain_id=$1 and proposal_id=$2 and voter=$3;

-- name: InsertIdxProcessLastBlock :exec
insert into idx_process (chain_id, process_name, last_block) values ($1, $2, $3);

-- name: UpdateIdxProcessLastBlock :exec
update idx_process set last_block=$3 where chain_id=$1 and process_name=$2;

-- name: InsertIdxSkippedBlock :exec
insert into idx_skipped_block (chain_id, process_name, skipped_block, reason) values ($1, $2, $3, $4);
