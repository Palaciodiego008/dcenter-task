package main

import (
	"docucenter-task/actions"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	// Crear el enrutador HTTP
	router := mux.NewRouter()

	// Definir las rutas de la API
	router.HandleFunc("/api/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the API!")
	}).Methods("GET")

	router.HandleFunc("/api/logistics/truck", actions.CreateTruckDelivery).Methods("POST")
	router.HandleFunc("/api/logistics/ship", actions.CreateShipDelivery).Methods("POST")
	router.HandleFunc("/api/logistics/deliveries", actions.GetDeliveries).Methods("GET")
	router.HandleFunc("/api/logistics/client/", actions.CreateClient).Methods("POST")

	// Iniciar el servidor HTTP
	fmt.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
