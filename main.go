package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kayceeDev/CRM-Project/services"
)

func main() {
	filename := "customers.json"
	const PORT = 8080
	router := mux.NewRouter().StrictSlash(true)
	fs := http.FileServer(http.Dir("./static"))
	router.Handle("/", http.StripPrefix("/", fs))
	router.HandleFunc("/customers", services.AddCustomer(filename)).Methods("POST")
	router.HandleFunc("/customers", services.GetAllCustomers(filename)).Methods(("GET"))
	router.HandleFunc("/customers/{id}", services.GetOneCustomer(filename)).Methods(("GET"))
	router.HandleFunc("/customers/{id}", services.UpdateCustomer(filename)).Methods(("PUT"))
	router.HandleFunc("/customers/{id}", services.DeleteCustomer(filename)).Methods(("DELETE"))
	fmt.Println("Starting server in Port....", PORT)
	log.Fatal(http.ListenAndServe(":8080", router))
}
