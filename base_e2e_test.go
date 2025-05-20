package main

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/k1LoW/runn"
)

func TestRouter(t *testing.T) {
	ctx := context.Background()
	// dsn := "username:password@tcp(localhost:3306)/testdb"
	// db, err := sql.Open("mysql", dsn)
	// dbr, err := sql.Open("mysql", dsn)

	ts := httptest.NewServer(NewRouter())
	t.Cleanup(func() {
		ts.Close()
	})
	opts := []runn.Option{
		runn.T(t),
		runn.Book("testdata/books/login.yml"),
		runn.Runner("req", ts.URL),
		// runn.DBRunner("db", dbr),
	}
	o, err := runn.New(opts...)
	if err != nil {
		t.Fatal(err)
	}
	if err := o.Run(ctx); err != nil {
		t.Fatal(err)
	}
}
