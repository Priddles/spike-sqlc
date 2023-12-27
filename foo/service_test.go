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

func TestService_ListFoos(t *testing.T) {
	dbConn, err := pgx.Connect(context.Background(), "postgres://admin:admin@localhost/db?sslmode=disable")
	require.NoError(t, err)
	defer dbConn.Close(context.Background())

	_, err = dbConn.Exec(context.Background(), `
		TRUNCATE TABLE foo CASCADE
	`)
	require.NoError(t, err)

	id := uuid.New()
	q := db.New(dbConn)
	_, err = q.CreateFoo(context.Background(), db.CreateFooParams{
		ID:       id,
		Location: ewkb.New(orb.Point{1, 1}, 4326),
	})
	require.NoError(t, err)

	type fields struct {
		dbSrv *db.Queries
	}
	type args struct {
		ctx    context.Context
		params db.ListFoosParams
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
		want    []db.ListFoosRow
	}{
		{
			name: "should not return geometry",
			fields: fields{
				dbSrv: db.New(dbConn),
			},
			args: args{
				ctx: context.Background(),
				params: db.ListFoosParams{
					ReturnGeometry: false,
					Bound:          ewkb.New(orb.MultiPoint{{0, 0}, {1, 1}}.Bound(), 4326),
				},
			},
			want: []db.ListFoosRow{
				{
					ID: id,
				},
			},
		},
		{
			name: "should return geometry",
			fields: fields{
				dbSrv: db.New(dbConn),
			},
			args: args{
				ctx: context.Background(),
				params: db.ListFoosParams{
					ReturnGeometry: true,
					Bound:          ewkb.New(orb.MultiPoint{{0, 0}, {1, 1}}.Bound(), 4326),
				},
			},
			want: []db.ListFoosRow{
				{
					ID:       id,
					Location: ewkb.New(orb.Point{1, 1}, 4326),
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f := &Service{
				dbSrv: test.fields.dbSrv,
			}
			got, err := f.ListFoos(test.args.ctx, test.args.params)
			assert.ErrorIs(t, err, test.wantErr)
			assert.Equal(t, test.want, got)
		})
	}
}
