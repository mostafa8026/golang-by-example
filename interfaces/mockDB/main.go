package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mostafa8026/golang-by-example/interfaces/mockDB/models"
)

// Env Global environment variables
type Env struct {
	db models.DataStore
}

func main() {
	db, err := models.NewDB("server=localhost\\sql2017;user id=sa;password=123456789!@#abc;database=bookstore")
	if err != nil {
		log.Panic(err)
	}

	env := &Env{db}

	http.HandleFunc("/books", env.BooksIndex)
	http.ListenAndServe(":3000", nil)
}

func (env *Env) BooksIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}
	bks, err := env.db.AllBooks()
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}
	for _, bk := range bks {
		fmt.Fprintf(w, "%v, %s, %s, $%.2f\n", bk.Id, bk.Title, bk.Author, bk.Price)
	}
}
