package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	r.Get("/vegetables", AllVegetables)

	http.ListenAndServe(":3333", r)
}
