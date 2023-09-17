-- name: InsertAuthUserWithEmail :exec
insert into auth_user (id, email, created_at) values (gen_random_uuid(), $1, now());

-- name: InsertAuthUserWithAddress :exec
insert into auth_user (id, address, created_at) values (gen_random_uuid(), $1, now());

-- name: InsertAuthUserWithUsername :exec
insert into auth_user (id, username, created_at) values (gen_random_uuid(), $1, now());

-- name: UpdateAuthUserEmail :exec
update auth_user set email=$2 where id=$1;

-- name: UpdateAuthUserAddress :exec
update auth_user set address=$2 where id=$1;

-- name: UpdateAuthUserUsername :exec
update auth_user set username=$2 where id=$1;

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
