package models

import (
	"database/sql"
	"errors"
)

type URL struct {
	ID       uint64
	ShortURL string
	Original string
}

type URLModel struct {
	DB *sql.DB
}

func (m *URLModel) Insert(shortURL, original string) (int64, error) {
	stmt := `INSERT INTO urls (short_url, original) VALUES (?, ?)`

	result, err := m.DB.Exec(stmt, shortURL, original)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *URLModel) Get(shortURL string) (*URL, error) {
	stmt := `SELECT id, short_url, original FROM urls WHERE short_url = ?`
	row := m.DB.QueryRow(stmt, shortURL)

	var url URL
	err := row.Scan(&url.ID, &url.ShortURL, &url.Original)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		}
	}

	return &url, nil
}
