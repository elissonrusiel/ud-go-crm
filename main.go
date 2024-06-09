package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

type Customer struct {
	ID        string
	Name      string
	Role      string
	Email     string
	Phone     string
	Contacted bool
}

var customers = map[string]Customer{
	"1": Customer{ID: "1", Name: "Alice", Role: "Software Engineer", Email: "alice@example.com", Phone: "6785556789", Contacted: false},
	"2": Customer{ID: "2", Name: "Bob", Role: "Account Manager", Email: "bob@example.com", Phone: "6785556780", Contacted: true},
	"3": Customer{ID: "3", Name: "Charles", Role: "Cybersecurity Sales Manager", Email: "charles@example.com", Phone: "6785556781", Contacted: true},
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customers)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["ID"]

	if _, ok := customers[id]; ok {
		customer := customers[id]
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customer)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newEntry Customer

	reqBody, _ := io.ReadAll(r.Body)
	json.Unmarshal(reqBody, &newEntry)

	if _, ok := customers[newEntry.ID]; ok {
		w.WriteHeader(http.StatusConflict)
	} else {
		customers[newEntry.ID] = newEntry
		w.WriteHeader(http.StatusCreated)
	}

	json.NewEncoder(w).Encode(customers)
}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newEntry Customer

	reqBody, _ := io.ReadAll(r.Body)
	json.Unmarshal(reqBody, &newEntry)

	if _, ok := customers[newEntry.ID]; ok {
		customers[newEntry.ID] = newEntry
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customers[newEntry.ID])
	} else {
		customers[newEntry.ID] = newEntry
		w.WriteHeader(http.StatusNotFound)
	}
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["ID"]

	if _, ok := customers[id]; ok {
		delete(customers, id)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customers)
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(customers)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFile(w, r, "./static/index.html")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/api/v1/customers", getCustomers).Methods("GET")
	router.HandleFunc("/api/v1/customers/{ID}", getCustomer).Methods("GET")
	router.HandleFunc("/api/v1/customers", addCustomer).Methods("POST")
	router.HandleFunc("/api/v1/customers/{ID}", updateCustomer).Methods("PUT")
	router.HandleFunc("/api/v1/customers/{ID}", deleteCustomer).Methods("DELETE")

	fmt.Println("Server is starting on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", router))
}
