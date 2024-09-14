package main

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type URLData struct {
	id        int
	long_url  string
	short_url string
	timestamp string
}

var db *sql.DB

func initDB() error {
	var err error
	db, err = sql.Open("sqlite3", "urlshort.db")
	return err
}

func closeDB() {
	db.Close()
}

func insertURLToDatabase(long string, short string) (string, error) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	_, err := db.Exec("INSERT INTO items (long_url, short_url, timestamp) VALUES (?, ?, ?)", long, short, currentTime)
	return currentTime, err
}

func getLongURLFromDatabase(short_url string) (string, error) {
	query := `SELECT long_url FROM items WHERE short_url = ?`
	rows := db.QueryRow(query, short_url)

	var original_url string
	original_url_not_found_message := original_url + " not found in database"
	err := rows.Scan(&original_url)
	if err != nil {
		if err == sql.ErrNoRows {
			return original_url_not_found_message, err
		}

		return original_url_not_found_message, err
	}

	return original_url, err
}

func getURLInfoFromDatabase(short_url string) (URLData, error) {
	query := `SELECT * FROM items WHERE short_url = ?`
	rows := db.QueryRow(query, short_url)

	var url_data URLData
	err := rows.Scan(&url_data.id, &url_data.long_url, &url_data.short_url, &url_data.timestamp)
	if err != nil {
		if err == sql.ErrNoRows {
			return url_data, err
		}

		return url_data, err
	}

	return url_data, err
}
