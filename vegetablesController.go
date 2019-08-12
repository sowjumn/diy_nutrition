package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Vegetable struct {
	ID       int
	Name     string
	Calories int
}

type VegetableResponse struct {
	Status     string
	Code       int
	Vegetables []Vegetable
}

func AllVegetables(w http.ResponseWriter, req *http.Request) {
	jsonVegetables := []Vegetable{
		{ID: 1, Name: "Potato", Calories: 130},
		{ID: 2, Name: "Broccoli", Calories: 50},
	}

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
