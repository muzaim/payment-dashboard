package db

import "database/sql"

func Migrate(db *sql.DB) error {
	stmts := []string{
		`CREATE TABLE IF NOT EXISTS users (
		  id INTEGER PRIMARY KEY AUTOINCREMENT,
		  email TEXT NOT NULL UNIQUE,
		  password_hash TEXT NOT NULL,
		  role TEXT NOT NULL
		);`,

		`CREATE TABLE IF NOT EXISTS payments (
		  id INTEGER PRIMARY KEY AUTOINCREMENT,
		  payment_id TEXT NOT NULL UNIQUE,
		  merchant TEXT NOT NULL,
		  status TEXT NOT NULL,
		  amount INTEGER NOT NULL,
		  created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);`,
	}

	for _, s := range stmts {
		if _, err := db.Exec(s); err != nil {
			return err
		}
	}

	return nil
}
