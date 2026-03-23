package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os" 
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := "127.0.0.1:3306"
	dbName := "mysql"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, pass, host, dbName)
	db, _ := sql.Open("mysql", dsn)

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		if err := db.Ping(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	http.ListenAndServe(":8080", nil)
}
