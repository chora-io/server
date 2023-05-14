-- name: PostData :exec
insert into data (iri, jsonld) values ($1, $2);
