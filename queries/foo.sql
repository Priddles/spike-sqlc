-- name: GetFoo :one
SELECT * FROM foo
WHERE id = $1
LIMIT 1;

-- name: CreateFoo :one
INSERT INTO foo (
    id, location
) VALUES (
    $1, ST_GeomFromEWKB(@location)
)
RETURNING *;
