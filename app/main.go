package main

import (
	"controllers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func handlerDirectory() http.Handler {
	return http.FileServer(http.Dir("./static"))
}

func main() {
	// Instantiate a new router by invoking the "NewRouter" handler
	router := mux.NewRouter().StrictSlash(true)

	// serve the HTML file to the home route
	router.Handle("/", handlerDirectory())

	// set handlers
	router.HandleFunc("/customers", controllers.GetCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", controllers.GetCustomer).Methods("GET")
	router.HandleFunc("/customers", controllers.CreateCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", controllers.DeleteCustomer).Methods("DELETE")
	router.HandleFunc("/customers/{id}", controllers.UpdateCustomer).Methods("PATCH")

	fmt.Println("Server is starting on port 3000...")
	http.ListenAndServe(":3000", router)
}
