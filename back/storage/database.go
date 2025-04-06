package storage

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Function return open database connection
func Open() (*sql.DB, error) { 

	db, err := sql.Open("sqlite3", "storage.db")
	if err != nil {
		return nil, err
	}

	if err := Init(db); err != nil {
		return nil, err
	}

	return db, nil

}

// Fuciton if table is not exist create in db sended as a argument 
func Init(db *sql.DB) error {
	schema := `
	CREATE TABLE IF NOT EXISTS Tables (
		id					  INTEGER PRIMARY KEY AUTOINCREMENT,
		title				  TEXT NOT NULL,
		contentHTMLpath		  TEXT NOT NULL,
		contentMarkdownPath   TEXT NOT NULL,
		data DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := db.Exec(schema)
	return err
}

// CRUD

// Create
func CreateTable(db *sql.DB, title string, contentPath string) (int64, error) {
	result, err := db.Exec("INSERT INTO Tables (title, contentPath) VALUES (?, ?)", title, contentPath)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Read
func GetTable(db *sql.DB, id int64) (Table, error) {
	var table Table
	err := db.QueryRow("SELECT id, title, contentPath, data FROM Tables WHERE id = ?", id).Scan(&table.ID, &table.Title, &table.ContentPath, &table.Data)
	if err != nil {
		return Table{}, err
	}
	return table, nil
}

func GetTables(db *sql.DB) ([]Table, error) {
	rows, err := db.Query("SELECT id, title, contentPath, data FROM Tables")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tables []Table
	for rows.Next() {
		var table Table
		err := rows.Scan(&table.ID, &table.Title, &table.ContentPath, &table.Data)
		if err != nil {
			return nil, err
		}
		tables = append(tables, table)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tables, nil
}

// Update
func UpdateTable(db *sql.DB, id int64, title string, contentPath string) error {
	_, err := db.Exec("UPDATE Tables SET title = ?, contentPath = ? WHERE id = ?", title, contentPath, id)
	return err
}

// Delete
func DeleteTable(db *sql.DB, id int64) error {
	_, err := db.Exec("DELETE FROM Tables WHERE id = ?", id)
	return err
}

type Table struct {
	ID          int64
	Title       string
	ContentPath string
	Data        string
}
