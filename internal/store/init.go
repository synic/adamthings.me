package store

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func Init(location string) (*Queries, error) {
	conn, err := sql.Open(
		"sqlite3",
		fmt.Sprintf("file:%s?cache=shared&mode=rwc&_journal_mode=WAL", location),
	)

	if err != nil {
		return nil, err
	}

	MaybeRunMigrations("sqlite3", conn)
	db := New(conn)
	return db, nil
}
