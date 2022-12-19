-- name: PostContent :one
INSERT INTO content (body)
VALUES ($1)
RETURNING (id);
