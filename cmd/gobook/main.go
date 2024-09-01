package main

import (
	"database/sql"
	"gobooks/internal/cli"
	"gobooks/internal/service"
	"gobooks/internal/web"
	"net/http"
	"os"

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

	if len(os.Args) > 1 && (os.Args[1] == "search" || os.Args[1] == "simulate") {
		booksCLI := cli.NewBookCLI(BookService)
		booksCLI.Run()
		return
	}

	router := http.NewServeMux()

	router.HandleFunc("GET /books", boooHandlers.GetBooks)
	router.HandleFunc("POST /books", boooHandlers.CreateBook)
	router.HandleFunc("GET /books/{id}", boooHandlers.GetBookByID)
	router.HandleFunc("PUT /books/{id}", boooHandlers.UpdateBook)
	router.HandleFunc("DELETE /books/{id}", boooHandlers.DeleteBook)
	router.HandleFunc("POST /books/simulate", boooHandlers.ReadBooks)

	http.ListenAndServe(":8080", router)

}
