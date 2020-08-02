package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Car struct
type Car struct {
	ID    string `json:id`
	Brand string `json:brand`
	Make  string `json:make`
}

var cars []Car

func getCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}
func getCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range cars {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var car Car
	_ = json.NewDecoder(r.Body).Decode(&car)
	car.ID = strconv.Itoa(rand.Intn(1000))
	cars = append(cars, car)
	json.NewEncoder(w).Encode(car)
}

func updateCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range cars {
		if item.ID == params["id"] {
			cars = append(cars[:index], cars[index+1:]...)
			var car Car
			_ = json.NewDecoder(r.Body).Decode(&car)
			car.ID = params["id"]
			cars = append(cars, car)
			json.NewEncoder(w).Encode(car)
			return
		}
	}
}
func deleteCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range cars {
		if item.ID == params["id"] {
			cars = append(cars[:index], cars[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(cars)
}

func main() {
	// iinit router
	r := mux.NewRouter()

	// mock data
	cars = append(cars, Car{ID: "1", Brand: "Tesla", Make: "Model S"})
	cars = append(cars, Car{ID: "2", Brand: "VW", Make: "Golf"})

	// routes
	r.HandleFunc("/api/cars", getCars).Methods("GET")
	r.HandleFunc("/api/car/{id}", getCar).Methods("GET")
	r.HandleFunc("/api/cars", createCar).Methods("POST")
	r.HandleFunc("/api/car/{id}", updateCar).Methods("PUT")
	r.HandleFunc("/api/car/{id}", deleteCar).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", r))
}
