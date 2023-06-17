-- name: GetData :one
select iri, jsonld from data where iri=$1;

-- name: GetIdxProcessLastBlock :one
select last_block from idx_process where chain_id=$1 and process_name=$2;
