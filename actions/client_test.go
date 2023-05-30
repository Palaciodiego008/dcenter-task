package actions

import (
	"bytes"
	"database/sql"
	"docucenter-task/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestCreateClientHandler(t *testing.T) {
	// Configurar la base de datos de prueba
	db, err := sql.Open("postgres", "postgres://user:password@localhost/testdb?sslmode=disable")
	assert.NoError(t, err)
	defer db.Close()

	// Crear el enrutador utilizando Gorilla Mux
	router := mux.NewRouter()

	// Asignar el enrutador a la función de manejo
	router.HandleFunc("/clients", CreateClient).Methods("POST")

	// Crear un servidor de prueba
	server := httptest.NewServer(router)
	defer server.Close()

	// Crear un cliente ficticio para enviar en la solicitud
	client := models.Client{
		ID:            23,
		Name:          "John",
		LastName:      "Doe",
		StreetAddress: "123 Main St.",
		Phone:         "1234567890",
		Email:         "jdoe@plp.com",
	}
	payload, err := json.Marshal(client)
	assert.NoError(t, err)

	// Crear una solicitud POST al endpoint "/clients" con el payload del cliente
	resp, err := http.Post(server.URL+"/clients", "application/json", bytes.NewBuffer(payload))
	assert.NoError(t, err)
	defer resp.Body.Close()

	// Verificar el código de estado de la respuesta
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	// Leer el cuerpo de la respuesta
	var createdClient models.Client
	err = json.NewDecoder(resp.Body).Decode(&createdClient)
	assert.NoError(t, err)

	// Verificar que el cliente fue creado correctamente
	assert.NotZero(t, createdClient.ID)
	assert.Equal(t, client.Name, createdClient.Name)
	assert.Equal(t, client.LastName, createdClient.LastName)
	assert.Equal(t, client.StreetAddress, createdClient.StreetAddress)
	assert.Equal(t, client.Phone, createdClient.Phone)
	assert.Equal(t, client.Email, createdClient.Email)
}
