package infrastructure

import "database/sql"

// make open conn sql.db postgres
func ConnectionPostgres() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	return db, err
}
