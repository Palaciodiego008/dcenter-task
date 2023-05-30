package actions

import (
	"docucenter-task/db"
	"docucenter-task/models"
	"encoding/json"
	"log"
	"net/http"
)

// Handler para consultar los planes de entrega con filtros
func GetDeliveries(w http.ResponseWriter, r *http.Request) {
	// Obtener los parámetros de consulta
	queryParams := r.URL.Query()

	// Crear una instancia de DB con la conexión establecida
	db, err := db.GetDBConnection()

	if err != nil {
		log.Fatal(err)
	}

	// Construir la consulta base
	query := db.Model(&models.TruckDelivery{})

	// Agregar filtros según los parámetros de consulta
	if clientID := queryParams.Get("client_id"); clientID != "" {
		query = query.Where("client_id = ?", clientID)
	}
	if productType := queryParams.Get("product_type"); productType != "" {
		query = query.Where("product_type = ?", productType)
	}
	// Agrega más filtros según sea necesario

	// Ejecutar la consulta
	var deliveries []models.TruckDelivery
	err = query.Find(&deliveries).Error
	if err != nil {
		log.Fatal(err)
	}

	// Convertir los resultados a JSON y enviar la respuesta
	response, err := json.Marshal(deliveries)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
