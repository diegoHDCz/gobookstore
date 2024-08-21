package main

import (
	"database/sql"
	"net/http"

	"github.com/diegoHDCz/gobooks/internal/service"
	"github.com/diegoHDCz/gobooks/internal/web"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	bookService := service.NewBookService(db)

	bookHandlers := web.NewBookHandlers(bookService)

	r := http.NewServeMux()
	r.HandleFunc("GET /books", bookHandlers.GetBooks)
	r.HandleFunc("POST /books", bookHandlers.CreateBook)
	r.HandleFunc("GET /books/{id}", bookHandlers.GetBookByID)
	r.HandleFunc("PUT /books/{id}", bookHandlers.UpdateBook)
	r.HandleFunc("DELETE /books/{ID}", bookHandlers.DeleteBook)

	http.ListenAndServe(":8080", r)
}
