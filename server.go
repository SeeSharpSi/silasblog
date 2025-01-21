package main

import (
	"context"
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
	port := flag.Int("port", 9779, "port the server runs on")
	address := flag.String("address", "localhost", "address the server runs on")
	flag.Parse()

	// ip parsing
	base_ip := *address
	ip := base_ip + ":" + strconv.Itoa(*port)
	// root_ip, err := url.Parse(ip)
	//if err != nil {
	//log.Panic(err)
	//}

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
}

func ServeStatic(w http.ResponseWriter, r *http.Request) {
	file := r.PathValue("file")
	log.Printf("got /static/%s request\n", file)
	http.ServeFile(w, r, "./static/"+file)
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
