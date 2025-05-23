package storage

import (
	"database/sql"
	"log"

	Article "plumlabs/back/articles/article"

	_ "github.com/mattn/go-sqlite3"
)

// Function return open database connection
func Open() (*sql.DB, error) {
	log.Printf("Openning database connection")

	db, err := sql.Open("sqlite3", "storage.db")
	if err != nil {
		return nil, err
	}

	if err := Init(db); err != nil {
		return nil, err
	}

	return db, nil

}

func Init(db *sql.DB) error {
	log.Printf("Init database")
	schema := `
	CREATE TABLE IF NOT EXISTS Articles (
    	id                    INTEGER PRIMARY KEY AUTOINCREMENT,
    	title                 TEXT NOT NULL UNIQUE,
    	htmlContent 		  TEXT NOT NULL,
    	mdContent			  TEXT NOT NULL,
    	created_at            DATETIME DEFAULT CURRENT_TIMESTAMP,
    	last_update           DATETIME DEFAULT CURRENT_TIMESTAMP 
	);
	`
	_, err := db.Exec(schema)
	return err
}

// CRUD

// CREATE
func CreateTable(db *sql.DB, article Article.Article) (int64, error) {

	log.Printf("Creating article with title: %s", article.Title)
	result, err := db.Exec("INSERT INTO Articles (title, mdContent,htmlContent) VALUES (?, ?, ?)", article.Title, article.MdContent, article.HtmlContent)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

// READ

func GetArticleById(db *sql.DB, id int) (*Article.Article, error) {
	log.Printf("Getting article with id: %d", id)
	var article Article.Article
	row := db.QueryRow("select id, title,mdContent,htmlContent,last_update from Articles WHERE id = ?", id)
	err := row.Scan(&article.Id, &article.Title, &article.MdContent, &article.HtmlContent)

	return &article, err
}
func GetArticleByTitle(db *sql.DB, title string) (*Article.Article, error) {
	log.Printf("Getting article with title: %s", title)
	var article Article.Article

	row := db.QueryRow("select id, title,mdContent,htmlContent,last_update from Articles WHERE title = ?", title)
	err := row.Scan(&article.Id, &article.Title, &article.MdContent, &article.HtmlContent, &article.LastUpdate)

	return &article, err
}

func GetAllArticles(db *sql.DB) ([]Article.Article, error) {
	log.Printf("Getting all articles")
	var _articles []Article.Article

	rows, err := db.Query("select id, title,mdContent,htmlContent,last_update from Articles")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var article Article.Article
		err := rows.Scan(&article.Id, &article.Title, &article.MdContent, &article.HtmlContent, &article.LastUpdate)
		if err != nil {
			log.Printf("Error scanning article row: %v", err)
			return nil, err
		}
		_articles = append(_articles, article)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating article rows: %v", err)
		return nil, err
	}

	return _articles, nil
}

// UPDATE

func UpdateAricle(db *sql.DB, a Article.Article) error {
	log.Printf("Updating article with id: %d", a.Id)
	_, err := db.Exec("UPDATE Articles set title = ?, mdContent = ?, htmlContent = ?, last_update = CURRENT_TIMESTAMP WHERE id = ?", a.Title, a.MdContent, a.HtmlContent, a.Id)
	if err != nil {
		return err
	}

	return nil
}

// DELETE
func DeleteArticle(db *sql.DB, title string) error {
	log.Printf("Deleting article with title: %s", title)
	_, err := db.Exec("DELETE FROM ARTICLES WHERE title = ?", title)
	if err != nil {
		return err
	}
	return nil
}
