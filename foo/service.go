package foo

import (
	"context"
	"foo/db"

	"github.com/google/uuid"
)

type Service struct {
	dbSrv *db.Queries
}

func New(dbSrv *db.Queries) (*Service, error) {
	return &Service{
		dbSrv: dbSrv,
	}, nil
}

func (f *Service) CreateFoo(ctx context.Context, params db.CreateFooParams) (db.Foo, error) {
	return f.dbSrv.CreateFoo(ctx, params)
}

func (f *Service) GetFoo(ctx context.Context, id uuid.UUID) (db.Foo, error) {
	return f.dbSrv.GetFoo(ctx, id)
}

func (f *Service) ListFoos(ctx context.Context, params db.ListFoosParams) ([]db.ListFoosRow, error) {
	return f.dbSrv.ListFoos(ctx, params)
}
