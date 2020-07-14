package main

import (
	// "fmt"
	"net/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello " + r.URL.Query().Get("name") + "!"))
	})

	http.ListenAndServe(":3000", r)

}