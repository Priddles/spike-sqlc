-- name: GetFoo :one
SELECT * FROM foo WHERE id = $1 LIMIT 1;

-- name: ListFoos :many
SELECT id,
    area,
    (CASE WHEN @return_geometry::boolean THEN location ELSE null END)::geom_point as location
FROM foo
WHERE location && @bound::geom_bound;

-- name: CreateFoo :one
INSERT INTO foo (id, location)
VALUES ($1, $2)
RETURNING *;
