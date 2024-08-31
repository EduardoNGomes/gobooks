package main

import (
	"database/sql"
	"gobooks/internal/service"
	"gobooks/internal/web"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	BookService := service.NewBookService(db)

	boooHandlers := web.NewBookhandlers(BookService)

	router := http.NewServeMux()

	router.HandleFunc("GET /books", boooHandlers.GetBooks)
	router.HandleFunc("POST /books", boooHandlers.CreateBook)
	router.HandleFunc("GET /books/{id}", boooHandlers.GetBookByID)
	router.HandleFunc("PUT /books/{id}", boooHandlers.UpdateBook)
	router.HandleFunc("DELETE /books/{id}", boooHandlers.DeleteBook)

	http.ListenAndServe(":8080", router)

}
