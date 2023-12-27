package foo

import (
	"context"
	"foo/db"
	"foo/ewkb"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/paulmach/orb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestService_CreateFoo(t *testing.T) {
	dbConn, err := pgx.Connect(context.Background(), "postgres://admin:admin@localhost/db?sslmode=disable")
	require.NoError(t, err)
	defer dbConn.Close(context.Background())

	id := uuid.New()

	type fields struct {
		dbSrv *db.Queries
	}
	type args struct {
		ctx    context.Context
		params db.CreateFooParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
		want    db.Foo
	}{
		{
			name: "should create foo",
			fields: fields{
				dbSrv: db.New(dbConn),
			},
			args: args{
				ctx: context.Background(),
				params: db.CreateFooParams{
					ID: id,
					Location: ewkb.Point{
						Geom:  orb.Point{1, 2},
						SRID:  4326,
						Valid: true,
					},
				},
			},
			want: db.Foo{
				ID: id,
				Location: ewkb.Point{
					Geom:  orb.Point{1, 2},
					SRID:  4326,
					Valid: true,
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f := &Service{
				dbSrv: test.fields.dbSrv,
			}
			got, err := f.CreateFoo(test.args.ctx, test.args.params)
			assert.ErrorIs(t, err, test.wantErr)
			assert.Equal(t, test.want, got)
		})
	}
}
