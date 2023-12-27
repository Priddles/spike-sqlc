// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: foo.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const getFoo = `-- name: GetFoo :one
SELECT id, area, location FROM foo WHERE id = $1 LIMIT 1
`

func (q *Queries) GetFoo(ctx context.Context, id uuid.UUID) (Foo, error) {
	row := q.db.QueryRow(ctx, getFoo, id)
	var i Foo
	err := row.Scan(&i.ID, &i.Area, &i.Location)
	return i, err
}
