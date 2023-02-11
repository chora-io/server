-- name: GetData :one
SELECT iri, jsonld
FROM data
WHERE iri=$1;
