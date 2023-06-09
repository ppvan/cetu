package models

import (
	"database/sql"
	"time"
)

type URL struct {
	ID          int64     `json:"id"`
	ShortURL    string    `json:"shortUrl"`
	OriginalURL string    `json:"originalUrl"`
	Clicks      int64     `json:"clicks"`
	ExpiredTime time.Time `json:"expiredTime"`
}

type URLModel struct {
	DB *sql.DB
}

// func (m *URLModel) Insert(shortURL, original string) (int64, error) {
// 	stmt := `INSERT INTO urls (short_url, original) VALUES (?, ?)`

// 	result, err := m.DB.Exec(stmt, shortURL, original)
// 	if err != nil {
// 		return 0, err
// 	}

// 	id, err := result.LastInsertId()
// 	if err != nil {
// 		return 0, err
// 	}

// 	return id, nil
// }

func (m *URLModel) GetLastInsertId() (int64, error) {
	stmt := `SELECT MAX(id) FROM urls;`
	row := m.DB.QueryRow(stmt)

	var lastID int64
	err := row.Scan(&lastID)
	if err != nil {
		return 0, err
	}

	return lastID, nil
}

// func (m *URLModel) Get(shortURL string) (*URL, error) {
// 	stmt := `SELECT id, short_url, original FROM urls WHERE short_url = ?`
// 	row := m.DB.QueryRow(stmt, shortURL)

// 	var url URL
// 	err := row.Scan(&url.ID, &url.ShortURL, &url.Original)
// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			return nil, ErrNoRecord
// 		}
// 	}

// 	return &url, nil
// }
