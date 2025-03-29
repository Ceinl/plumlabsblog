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
	CREATE TABLE IF NOT EXISTS Tables (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		contentPath TEXT NOT NULL,
		data DATETIME DEFAULT CURRENT_TIMESTAMP,
	);`
	
	_, err := db.Exec(schema)
	return err
} 

func InsertTable(db *sql.DB, title string, contentPath string) error {
	sqlStmt := ` INSERT INTO tables(title, conentPath) VALUES (?,?) ; `
	
	_, err := db.Exec(sqlStmt, title, contentPath)
	return err
}

func GetTables(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SELECT * FROM tables")
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
