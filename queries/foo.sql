-- name: GetFoo :one
SELECT * FROM foo WHERE id = $1 LIMIT 1;
