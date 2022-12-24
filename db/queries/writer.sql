-- name: PostData :one
INSERT INTO data (canon, context, jsonld)
VALUES ($1, $2, $3)
RETURNING (id);
