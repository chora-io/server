-- name: GetData :one
SELECT iri, context, jsonld
FROM data
WHERE iri=$1;
