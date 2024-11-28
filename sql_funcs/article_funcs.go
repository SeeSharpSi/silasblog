package sql_funcs

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Article struct {
	Id      int
	Title   string
	Teaser  string
	Content string
	Clicks  int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func connect_db() *sql.DB {
	db, err := sql.Open("sqlite3", "./data.db")
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

func GetArticle(id int) (article Article, e error) {
	db := connect_db()
	defer db.Close()
	sqlStmt := fmt.Sprintf("SELECT article_id, title, content FROM articles WHERE article_id = %d", id)
	result, err := db.Query(sqlStmt)
	fmt.Printf("%+v\n", result)
	defer result.Close()
	check(err)
	result.Next()
	if _, err := result.Columns(); err != nil {
		return Article{}, errors.New(fmt.Sprintf("Article with id %d does not exist, exiting with error: %v", id, err))
	}
	result.Scan(&article.Id, &article.Title, &article.Content)
	return article, nil
}

func AddClick(id int) {
	db := connect_db()
	defer db.Close()
	sqlStmt := fmt.Sprintf("UPDATE articles SET clicks = clicks + 1 WHERE article_id = %d", id)
	_, err := db.Exec(sqlStmt)
	check(err)
}
