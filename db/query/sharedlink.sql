-- name: CreateSharedlink :one
INSERT INTO sharedlinks (
  name,
  urlhash
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetSharedlink :one
SELECT * FROM sharedlinks
WHERE id = $1 LIMIT 1;

-- name: ListSharedlink :many
SELECT * FROM sharedlinks
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateSharedlink :one
UPDATE sharedlinks
SET name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteSharedlink :exec
DELETE FROM sharedlinks
WHERE id = $1;