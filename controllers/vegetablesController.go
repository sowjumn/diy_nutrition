package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/sowjumn/diy_nutrition/models"
)

type Vegetable struct {
	ID       int
	Name     string
	Calories int
}

type VegetableResponse struct {
	Status     string
	Code       int
	Vegetables []*models.VegetableRecord
}

// func AllVegetables(w http.ResponseWriter, req *http.Request) {
// 	jsonVegetables := []models.VegetableRecord{
// 		{ID: 1, Name: "Potato", Calories: 130},
// 		{ID: 2, Name: "Broccoli", Calories: 50},
// 	}

// 	vr := VegetableResponse{
// 		Status:     "ok",
// 		Code:       200,
// 		Vegetables: jsonVegetables,
// 	}

// 	resp, err := json.Marshal(vr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(200)
// 	w.Write(resp)
// }

func GetVegetable(w http.ResponseWriter, req *http.Request) {
	fmt.Printf(chi.URLParam(req, "id"))
	// jsonVegetables := []Vegetable{
	// 	{ID: 1, Name: "Potato", Calories: 130},
	// }
	id, _ := strconv.Atoi(chi.URLParam(req, "id"))
	jsonVegetables := models.GetRecord(id)
	vr := VegetableResponse{
		Status:     "ok",
		Code:       200,
		Vegetables: jsonVegetables,
	}

	resp, err := json.Marshal(vr)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(resp)
}

func AddVegetable(w http.ResponseWriter, req *http.Request) {

}

func UpdateVegetable(w http.ResponseWriter, req *http.Request) {

}

func DeleteVegetable(w http.ResponseWriter, req *http.Request) {

}
