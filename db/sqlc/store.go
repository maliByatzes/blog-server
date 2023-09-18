package db

import "database/sql"

// Store defines all functions to execute database queries and transactiosns
type Store interface {
	Querier
}

// SQLStore provides all functions to execute database queries and transactions
type SQLStore struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new store
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}
