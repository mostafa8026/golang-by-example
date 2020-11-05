package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mostafa8026/golang-by-example/interfaces/mockDB/models"
)

type mockDB struct{}

func (db *mockDB) AllBooks() ([]*models.Book, error) {
	bks := make([]*models.Book, 0)
	bks = append(bks, &models.Book{1, "test", "testAuthor", 100})
	bks = append(bks, &models.Book{2, "test2", "testAuthor2", 100})
	return bks, nil
}

func TestBooksIndex(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/books", nil)

	env := &Env{db: &mockDB{}}
	http.HandlerFunc(env.BooksIndex).ServeHTTP(rec, req)

	expected := "1, test, testAuthor, $100.00\n2, test2, testAuthor2, $100.00\n"
	if expected != rec.Body.String() {
		t.Errorf("")
	}
}
