package psql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Database struct {
	*pgx.Conn
}

func Connect(dns string) (*Database, error) {
	conn, err := pgx.Connect(context.Background(), dns)
	if err != nil {
		return nil, err
	}

	return &Database{conn}, nil
}

func (d *Database) Migrate() error {
	_, err := d.Exec(context.Background(), `
		CREATE TABLE IF NOT EXISTS books (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			title TEXT NOT NULL,
			author TEXT NOT NULL,
			description TEXT NOT NULL
		);
	`)
	if err != nil {
		return err
	}

	return nil
}
