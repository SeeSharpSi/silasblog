package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"

	sql "seesharpsi/silasblog/sql_funcs"
	"seesharpsi/silasblog/templ"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Define a command-line flag for the database path.
	// The default value is "/data.db", which is suitable for the Docker container.
	dbPath := flag.String("db", "/data.db", "path to the SQLite database file")
	flag.Parse()

	// Set the database path for the sql_funcs package.
	sql.SetDBPath(*dbPath)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9779"
	}

	// For Cloud Run, we listen on all available interfaces, signified by a leading ":"
	ip := ":" + port

	mux := http.NewServeMux()
	add_routes(mux)

	server := http.Server{
		Addr:    ip,
		Handler: mux,
	}

	// start server
	log.Printf("running server on %s\n", ip)
	err := server.ListenAndServe()
	defer server.Close()
	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("server closed\n")
	} else if err != nil {
		log.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func add_routes(mux *http.ServeMux) {
	mux.HandleFunc("/", GetIndex)
	mux.HandleFunc("/static/{file}", ServeStatic)
	mux.HandleFunc("/articles", GetArticles)
	mux.HandleFunc("/articlethumbnail/{id}", GetArticleThumbnail)
	mux.HandleFunc("/article/{id}", GetArticle)
	mux.HandleFunc("/api/all", GetAllData)
	mux.HandleFunc("/game/questions", GetGame)
	mux.HandleFunc("/game/new-question", GetNewQuestion)
}

func GetGame(w http.ResponseWriter, r *http.Request) {
	log.Printf("got /game/questions request\n")
	question := sql.GetRandomQuestion()
	component := templ.Game(question)
	component.Render(context.Background(), w)
}

func GetNewQuestion(w http.ResponseWriter, r *http.Request) {
	log.Printf("got /game/new-question request\n")
	question := sql.GetRandomQuestion()
	component := templ.Question(question)
	component.Render(context.Background(), w)
}



func GetAllData(w http.ResponseWriter, r *http.Request) {
	log.Printf("got /api/all request\n")
	articles, err := sql.GetAllArticles()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("error getting all articles: %s\n", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

func ServeStatic(w http.ResponseWriter, r *http.Request) {
	file := r.PathValue("file")
	log.Printf("got /static/%s request\n", file)
	http.ServeFile(w, r, "static/"+file)
}

func GetIndex(w http.ResponseWriter, r *http.Request) {
	log.Printf("got / request\n")
	component := templ.Index()
	component.Render(context.Background(), w)
}

func GetArticles(w http.ResponseWriter, r *http.Request) {
	log.Printf("got /articles request\n")
	component := templ.Articles()
	component.Render(context.Background(), w)
}

func GetArticleThumbnail(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	log.Printf("got /articlethumbnail/%s request\n", id)
	int_id, err := strconv.Atoi(id)
	check(err)
	cont := true
	articlethumb, err := sql.GetArticleThumbnail(int_id)
	if err != nil {
		cont = false
	}
	component := templ.ArticleThumbnail(articlethumb, cont)
	component.Render(context.Background(), w)
}

func GetArticle(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	log.Printf("got /article/%s request\n", id)
	int_id, err := strconv.Atoi(id)
	check(err)
	article, err := sql.GetArticle(int_id)
	check(err)
	component := templ.Article(article)
	component.Render(context.Background(), w)
}
