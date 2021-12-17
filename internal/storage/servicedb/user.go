package servicedb

import (
	"context"

	"github.com/sundaytycoon/buttons-api/edge/mysql"
	constantsmodel "github.com/sundaytycoon/buttons-api/internal/constants/model"
	constantsquery "github.com/sundaytycoon/buttons-api/internal/constants/query"
	"github.com/sundaytycoon/buttons-api/pkg/er"
)

type Storage struct{}

func New() *Storage {
	return &Storage{}
}

func (a *Storage) GetUser(ctx context.Context, tx mysql.ContextExecutor, id string) (*constantsmodel.User, error) {
	op := er.GetOperator()

	rows, err := tx.QueryContext(ctx, constantsquery.GetUser, id)
	if err != nil {
		return nil, er.WrapOp(err, op)
	}
	if rows.Err() != nil {
		return nil, er.WrapOp(rows.Err(), op)
	}

	var u = new(constantsmodel.User)

	for rows.Next() {
		rows.Scan(&u.ID, &u.Name, &u.State)
	}
	defer rows.Close()

	return u, nil
}
