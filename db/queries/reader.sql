-- name: GetData :one
SELECT id, canon, context, jsonld
FROM data
WHERE id=$1;
