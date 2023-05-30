package actions

import (
	"docucenter-task/db"
	"docucenter-task/models"
	"encoding/json"
	"fmt"
	"net/http"
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
