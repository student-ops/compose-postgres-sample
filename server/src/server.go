package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var host string

const (
	port     = 5432
	user     = "your-username"
	password = "your-password"
	dbname   = "your-dbname"
)

var db *sql.DB

func main() {
	host, ok := os.LookupEnv("DB_HOST")
	if !ok {
		fmt.Println("DB_HOST not set")
		return
	}
	var err error
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/", hello)
	http.HandleFunc("/db", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func handler(w http.ResponseWriter, r *http.Request) {
	count, err := getCount()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	fmt.Fprintf(w, "Access count: %d", count)

	_, err = db.Exec("UPDATE access_count SET count = count + 1 WHERE id = 1")
	if err != nil {
		log.Fatal(err)
	}
}

func getCount() (int, error) {
	var count int
	err := db.QueryRow("SELECT count FROM access_count WHERE id = 1").Scan(&count)
	if err != nil {
		if err == sql.ErrNoRows {
			_, err = db.Exec("INSERT INTO access_count (id, count) VALUES (1, 0)")
			if err != nil {
				return 0, err
			}
			return 0, nil
		}
		return 0, err
	}
	return count, nil
}
