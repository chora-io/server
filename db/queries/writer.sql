-- name: PostData :exec
INSERT INTO data (iri, context, jsonld)
VALUES ($1, $2, $3);
