package sql_funcs

import (
	"log"

	_ "modernc.org/sqlite"
)

type Question struct {
	ID   int
	Text string
}

func GetRandomQuestion() string {
	db := connect_db()
	defer db.Close()

	var question string
	err := db.QueryRow("SELECT text FROM questions ORDER BY RANDOM() LIMIT 1").Scan(&question)
	if err != nil {
		log.Fatal(err)
	}
	return question
}
