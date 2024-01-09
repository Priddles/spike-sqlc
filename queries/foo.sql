-- name: GetFoo :one
SELECT * FROM foo WHERE id = $1 LIMIT 1;

-- name: ListFoos :many
--
-- The use of the CASE statement to conditionally select the spatial columns confuses sqlc, so type
-- casting is used to inform which Go types should be used in the generated code. Without the type
-- casts, sqlc falls back to interface{}.
SELECT id,
    (CASE WHEN @return_geometry::boolean THEN area ELSE null END)::geom_any as area,
    (CASE WHEN @return_geometry::boolean THEN location ELSE null END)::geom_point as location
FROM foo
WHERE location && @bound::geom_bound;

-- name: CreateFoo :one
INSERT INTO foo (id, location)
VALUES ($1, $2)
RETURNING *;
