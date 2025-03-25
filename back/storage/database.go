package storage

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Open() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./storage.db")
	if err != nil {
		return nil, err
	}
	
	if err = db.Ping(); err != nil {
		return nil, err
	}
	
	return db, nil
}

func Init(db *sql.DB) error {
	schema := `
	CREATE TABLE IF NOT EXISTS tables (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		header TEXT DEFAULT '',
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`
	
	_, err := db.Exec(schema)
	return err
} 

func InsertTable(db *sql.DB, name string) error {
	sqlStmt := ` INSERT INTO tables(name) VALUES (?); `
	
	_, err := db.Exec(sqlStmt, name)
	return err
}

func GetTables(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SELECT name FROM tables")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tables []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		tables = append(tables, name)
	}

	return tables, nil
}
