-- name: GetContent :one
SELECT (body)
FROM content
WHERE (id=$1);
