package sql_funcs

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	_ "modernc.org/sqlite" // Pure Go SQLite driver
)

// dbPath holds the path to the SQLite database file.
// It's a package-level variable that can be set at runtime.
var dbPath string

// SetDBPath sets the path for the database file. This function should be called
// once during application initialization.
func SetDBPath(path string) {
	dbPath = path
	log.Printf("Database path set to: %s", dbPath)
}

type Article struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Teaser    string    `json:"teaser"`
	Content   string    `json:"content"`
	Clicks    int       `json:"clicks"`
	CreatedAt time.Time `json:"created_at"`
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func connect_db() *sql.DB {
	// Check if the dbPath has been set.
	if dbPath == "" {
		// Panicking here because the database path is a critical, required configuration.
		// The application cannot function without it.
		panic("database path is not set. Please call SetDBPath() before making database calls.")
	}
	// Use "sqlite" as the driver name for modernc.org/sqlite
	db, err := sql.Open("sqlite", dbPath)
	check(err)
	return db
}

func GetArticleThumbnail(id int) (Article, error) {
	articlethumb := Article{}
	db := connect_db()
	defer db.Close()
	sqlStmt := "SELECT article_id, title, teaser FROM articles WHERE article_id = ?"
	result := db.QueryRow(sqlStmt, id)
	err := result.Scan(&articlethumb.Id, &articlethumb.Title, &articlethumb.Teaser)
	if err == sql.ErrNoRows {
		return Article{}, err
	}
	return articlethumb, nil
}

func GetArticles() (articles []Article) {
	db := connect_db()
	defer db.Close()
	sqlStmt := fmt.Sprintf("SELECT article_id, title FROM articles;")
	result, err := db.Query(sqlStmt)
	defer result.Close()
	check(err)
	for result.Next() {
		article := Article{}
		result.Scan(&article.Id, &article.Title)
		articles = append(articles, article)
	}
	return articles
}

func GetAllArticles() ([]Article, error) {
	db := connect_db()
	defer db.Close()
	sqlStmt := "SELECT article_id, title, teaser, content, clicks, created_at FROM articles;"
	rows, err := db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []Article
	for rows.Next() {
		var article Article
		err := rows.Scan(&article.Id, &article.Title, &article.Teaser, &article.Content, &article.Clicks, &article.CreatedAt)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	return articles, nil
}

func GetArticle(id int) (article Article, e error) {
	db := connect_db()
	defer db.Close()
	sqlStmt := "SELECT article_id, title, content, created_at FROM articles WHERE article_id = ?"
	result, err := db.Query(sqlStmt, id)
	defer result.Close()
	check(err)
	result.Next()
	if _, err := result.Columns(); err != nil {
		return Article{}, errors.New(fmt.Sprintf("Article with id %d does not exist, exiting with error: %v", id, err))
	}

	var createdAtStr string
	err = result.Scan(&article.Id, &article.Title, &article.Content, &createdAtStr)
	if err != nil {
		return Article{}, err
	}

	// Parse the created_at string into time.Time
	if createdAtStr != "" {
		parsedTime, err := time.Parse("2006-01-02T15:04:05.999999999Z07:00", createdAtStr)
		if err != nil {
			// Try alternative formats
			parsedTime, err = time.Parse("2006-01-02T15:04:05Z07:00", createdAtStr)
			if err != nil {
				parsedTime, err = time.Parse("2006-01-02 15:04:05", createdAtStr)
				if err != nil {
					// Set to current time if parsing fails
					parsedTime = time.Now()
				}
			}
		}
		article.CreatedAt = parsedTime
	} else {
		article.CreatedAt = time.Now()
	}
	return article, nil
}

func AddClick(id int) {
	db := connect_db()
	defer db.Close()
	sqlStmt := "UPDATE articles SET clicks = clicks + 1 WHERE article_id = ?"
	_, err := db.Exec(sqlStmt, id)
	check(err)
}
