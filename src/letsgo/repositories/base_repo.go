package repositories

import "database/sql"

type BaseRepo struct {
	db *sql.DB
}

func NewBaseRepo(db *sql.DB) *BaseRepo {
	return &BaseRepo{
		db: db,
	}
}
