package models

import "database/sql"

func DatabaseStatusCheck(db *sql.DB) (string, error) {
	err := db.Ping()

	if err != nil {
		return "unavailable", err
	}

	return "available", nil
}
