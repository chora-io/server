-- name: PostData :exec
INSERT INTO data (iri, jsonld)
VALUES ($1, $2);
