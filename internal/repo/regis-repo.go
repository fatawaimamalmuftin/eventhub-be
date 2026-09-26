package repo

import (
	"context"

	cuserror "github.com/fatawaimamalmuftin/eventhub-be/internal/CusError"
	"github.com/fatawaimamalmuftin/eventhub-be/internal/dto"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DbRegisRepo struct {
	db *pgxpool.Pool
}

func RegisRepo(db *pgxpool.Pool) *DbRegisRepo {
	return &DbRegisRepo{
		db: db,
	}
}

func (d *DbRegisRepo) NewAccount(c context.Context, newUser dto.Regis) error {
	q := "INSERT INTO users (full_name,email,password) VALUES ($1, $2, $3);"
	arg := []any{newUser.FullName, newUser.Email, newUser.Password}

	data, err := d.db.Exec(c, q, arg...)

	if err != nil {
		return err
	}

	if data.RowsAffected() == 0 {
		return cuserror.ErrNoRowAffected
	}

	return nil
}
