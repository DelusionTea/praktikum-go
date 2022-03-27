package DataBase

import (
	"context"
	"database/sql"
	"fmt"
)

type PGDataBase struct {
	conn    *sql.DB
	baseURL string
}

func NewDatabase(baseURL string, db *sql.DB) *PGDataBase {
	result := &PGDataBase{
		conn:    db,
		baseURL: baseURL,
	}
	return result
}

func (db *PGDataBase) Ping(ctx context.Context) error {

	err := db.conn.PingContext(ctx)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
