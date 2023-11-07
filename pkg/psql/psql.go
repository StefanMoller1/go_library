package psql

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

type Database struct {
	*pgx.Conn
}

func Connect(dns string, log *log.Logger) (*Database, error) {
	conn, err := pgx.Connect(context.Background(), dns)
	if err != nil {
		return nil, err
	}

	log.Printf("connected to database: %s", dns)
	return &Database{conn}, nil
}

func (d *Database) Migrate(log *log.Logger) error {
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

	log.Println("created books database table")

	return nil
}
