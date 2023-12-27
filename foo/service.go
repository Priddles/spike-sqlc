package foo

import (
	"context"
	"foo/db"

	"github.com/google/uuid"
	"github.com/paulmach/orb"
)

type dbSrv interface {
	CreateFoo(context.Context, db.CreateFooParams) (db.Foo, error)
	GetFoo(context.Context, uuid.UUID) (db.Foo, error)
	ListFoos(context.Context, any) ([]db.Foo, error)
}

type Service struct {
	dbSrv dbSrv
}

func New(dbSrv dbSrv) (*Service, error) {
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

func (f *Service) ListFoos(ctx context.Context, bound orb.Bound) ([]db.Foo, error) {
	return f.dbSrv.ListFoos(ctx, bound)
}
