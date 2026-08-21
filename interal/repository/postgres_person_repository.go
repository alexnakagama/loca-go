package repository

import "database/sql"

type PostgresPersonRepository struct {
	db *sql.DB
}
