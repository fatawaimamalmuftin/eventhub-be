package repo

import "github.com/jackc/pgx/v5/pgxpool"

type DbLoginRepo struct {
	Db *pgxpool.Pool
}

func LoginRepo(db *pgxpool.Pool) *DbLoginRepo {
	return &DbLoginRepo{
		Db: db,
	}
}
