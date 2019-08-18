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

type VegetableResponse struct {
	Status     string
	Code       int
	Vegetables []models.VegetableRecord
}

func AllVegetables(w http.ResponseWriter, req *http.Request) {
	auth := checkAuth(w, req)
	if auth == false {
		return
	}

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

func checkAuth(w http.ResponseWriter, req *http.Request) bool {
	accessToken := "yes"
	_, p, _ := req.BasicAuth()
	auth := true
	if p != accessToken {
		w.WriteHeader(http.StatusUnauthorized)
		auth = false
	}
	return auth
}

func GetVegetable(w http.ResponseWriter, req *http.Request) {
	auth := checkAuth(w, req)
	if auth == false {
		return
	}

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

type vegetableInfo struct {
	name     string
	calories int
}

func AddVegetable(w http.ResponseWriter, req *http.Request) {
	auth := checkAuth(w, req)
	if auth == false {
		return
	}

	decoder := json.NewDecoder(req.Body)
	var vegpost vegetableInfo
	err := decoder.Decode(&vegpost)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Name: %v, Calories: %v", name, calories)
	models.AddRecord(name, calories)
}

func UpdateVegetable(w http.ResponseWriter, req *http.Request) {
	auth := checkAuth(w, req)
	if auth == false {
		return
	}

	decoder := json.NewDecoder(req.Body)
	var vegpost vegetableInfo
	err := decoder.Decode(&vegpost)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Name: %v, Calories: %v", id, name, calories)
	models.UpdateRecord(name, calories)
}

func DeleteVegetable(w http.ResponseWriter, req *http.Request) {
	auth := checkAuth(w, req)
	if auth == false {
		return
	}
}
