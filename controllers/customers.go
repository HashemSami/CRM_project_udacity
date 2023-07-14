package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"models"
	"net/http"

	"github.com/gorilla/mux"
)

func GetCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(models.MockData)
}

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	if _, ok := models.MockData[id]; ok {
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(models.MockData[id])

	} else {
		w.WriteHeader(http.StatusNotFound)
		Massage := map[string]string{
			"Err": "entry not found",
		}

		json.NewEncoder(w).Encode(Massage)
	}

}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// 1. get new entry
	var newEntry models.Customer

	// 2. read the request
	reqBody, _ := io.ReadAll(r.Body)

	// 3. parse JSON body
	json.Unmarshal(reqBody, &newEntry)

	// 4. add new entry to dictionary
	if _, ok := models.MockData[newEntry.ID]; ok {
		w.WriteHeader(http.StatusConflict)
		Massage := map[string]string{
			"Err": "entry already exist",
		}
		json.NewEncoder(w).Encode(Massage)

	} else {
		models.MockData[newEntry.ID] = newEntry
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newEntry)
	}

}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	if id != "" {
		fmt.Println("id", id)
	}
	if _, ok := models.MockData[id]; ok {
		delete(models.MockData, id)

		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(models.MockData)

	} else {
		w.WriteHeader(http.StatusNotFound)
		Massage := map[string]string{
			"Err": "entry not found",
		}

		json.NewEncoder(w).Encode(Massage)
	}
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]

	if data, ok := models.MockData[id]; ok {
		var newEntry models.Customer

		reqBody, _ := io.ReadAll(r.Body)
		json.Unmarshal(reqBody, &newEntry)

		if newEntry.Name != "" {
			data.Name = newEntry.Name
		}
		if newEntry.Email != "" {
			data.Email = newEntry.Email
		}
		if newEntry.Role != "" {
			data.Role = newEntry.Role
		}
		if newEntry.Phone != 0 {
			data.Phone = newEntry.Phone
		}
		if newEntry.Contacted != nil {
			data.Contacted = newEntry.Contacted
		}

		models.MockData[id] = data
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(models.MockData)

	} else {
		w.WriteHeader(http.StatusNotFound)
		Massage := map[string]string{
			"Err": "entry not found",
		}

		json.NewEncoder(w).Encode(Massage)
	}
}
