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
	Vegetables []models.VegetableRecord
}

func AllVegetables(w http.ResponseWriter, req *http.Request) {
	jsonVegetables, _ := models.GetAllRecords()

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

func GetVegetable(w http.ResponseWriter, req *http.Request) {
	status := "ok"
	code := 200
	id, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil {
		status = "Vegetable ID not an integer"
		code = 422
	}
	jsonVegetables, err := models.GetRecord(id)

	if err != nil {
		fmt.Printf("%v", err)
		status = "Cant find Vegetable"
		code = 404
	}

	vr := VegetableResponse{
		Status:     status,
		Code:       code,
		Vegetables: jsonVegetables,
	}

	resp, err := json.Marshal(vr)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(resp)
}

func AddVegetable(w http.ResponseWriter, req *http.Request) {

}

func UpdateVegetable(w http.ResponseWriter, req *http.Request) {

}

func DeleteVegetable(w http.ResponseWriter, req *http.Request) {

}
