package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) GetDb() *sql.DB{
	return d.db
}

func NewDatabase() (*Database, error) {
	db, err := sql.Open("postgres","postgresql://root:password@localhost:5433/cardgame_db?sslmode=disable")
	if err != nil {
		return nil, err
	}
	return &Database{db:db},nil
}
