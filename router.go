package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sowjumn/diy_nutrition/controllers"
)

func main() {
	r := chi.NewRouter()
	r.Get("/vegetables", controllers.AllVegetables)
	r.Get("/vegetables/{id}", controllers.GetVegetable)
	r.Post("/vegetables", controllers.AddVegetable)
	r.Put("/vegetables/{id}", controllers.UpdateVegetable)
	r.Delete("/vegetables/{id}", controllers.DeleteVegetable)
	http.ListenAndServe(":3333", r)
}
