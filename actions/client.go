package actions

import (
	"docucenter-task/db"
	"docucenter-task/models"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// ...

func CreateClient(w http.ResponseWriter, r *http.Request) {
	var client models.Client
	err := json.NewDecoder(r.Body).Decode(&client)
	if err != nil {
		http.Error(w, "Error al parsear los datos del cliente", http.StatusBadRequest)
		return
	}

	fmt.Println("client: ", client)

	db, err := db.GetDBConnection()
	if err != nil {
		http.Error(w, "Error al conectar a la base de datos", http.StatusInternalServerError)
		return
	}

	// Verificar si la tabla "clients" existe, y si no, crearla
	if !db.Migrator().HasTable(&models.Client{}) {
		err := db.Migrator().CreateTable(&models.Client{})
		if err != nil {
			http.Error(w, "Error al crear la tabla 'clients'", http.StatusInternalServerError)
			return
		}
	}

	// Insertar los datos en la tabla "clients"
	result := db.Create(&client)
	if result.Error != nil {
		http.Error(w, fmt.Sprintf("Error al insertar el cliente: %s", result.Error.Error()), http.StatusInternalServerError)
		return
	}

	fmt.Println("client: ", client)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(client)
}

func UpdateClient(w http.ResponseWriter, r *http.Request) {
	// Get the client ID from the request URL or body
	clientID := mux.Vars(r)["id"]
	if clientID == "" {
		http.Error(w, "Client ID not provided", http.StatusBadRequest)
		return
	}

	var updatedClient models.Client
	err := json.NewDecoder(r.Body).Decode(&updatedClient)
	if err != nil {
		http.Error(w, "Error decoding updated client data", http.StatusBadRequest)
		return
	}

	db, err := db.GetDBConnection()
	if err != nil {
		http.Error(w, "Error connecting to the database", http.StatusInternalServerError)
		return
	}

	// Check if the "clients" table exists, and if not, create it
	if !db.Migrator().HasTable(&models.Client{}) {
		err := db.Migrator().CreateTable(&models.Client{})
		if err != nil {
			http.Error(w, "Error creating 'clients' table", http.StatusInternalServerError)
			return
		}
	}

	// Check if the client exists
	var existingClient models.Client
	result := db.First(&existingClient, "id = ?", clientID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "Client not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error retrieving client data", http.StatusInternalServerError)
		}
		return
	}

	// Update the client data
	existingClient.Name = updatedClient.Name
	existingClient.Email = updatedClient.Email
	// Update other fields as needed

	// Save the updated client record
	result = db.Save(&existingClient)
	if result.Error != nil {
		http.Error(w, "Error updating client data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(existingClient)
}
