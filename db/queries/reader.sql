-- name: GetData :one
select iri, jsonld from data where iri=$1;
