package query

import (
	"context"
	"database/sql"
)

type Transaction interface {
	ExecuteTx(tx *sql.Tx) (sql.Result, error)
}

func ExecuteQueries(queries *[]Transaction, db *sql.DB) error {
	var con = context.Background()
	var tx, errT = db.BeginTx(con, nil)
	if errT != nil {
		var _ = db.Close()
		return errT
	}
	for _, q := range *queries {
		var _, err = q.ExecuteTx(tx)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}
