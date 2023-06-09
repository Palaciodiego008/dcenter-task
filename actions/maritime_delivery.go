package actions

import (
	"docucenter-task/db"
	"docucenter-task/internal"
	"docucenter-task/models"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// Handler para crear un plan de entrega de logística terrestre

// ...

func CreateTruckDelivery(w http.ResponseWriter, r *http.Request) {
	// Obtener el token de autenticación de la cabecera de autorización
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Token de autenticación requerido")
		return
	}

	// Validar el token
	validTk := internal.ValidateToken(tokenString)

	if !validTk {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Token de autenticación inválido")
		return
	}

	// Parsear los datos de la solicitud a una estructura TruckDelivery
	var truckDelivery models.TruckDelivery
	err := json.NewDecoder(r.Body).Decode(&truckDelivery)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Error al parsear los datos de entrada")
		return
	}

	validate := validator.New()
	err = validate.Struct(truckDelivery)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make([]string, 0, len(validationErrors))

		// Obtener los mensajes de error de validación
		for _, e := range validationErrors {
			errorMessages = append(errorMessages, fmt.Sprintf("Campo '%s' %s", e.Field(), e.Tag()))
		}

		// Devolver una respuesta de error con los mensajes de error
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Errores de validación: %v", errorMessages)
		return
	}

	// Calcular el descuento
	if truckDelivery.Quantity > 10 {
		discount := truckDelivery.ShippingPrice * 0.05
		truckDelivery.DiscountedPrice = truckDelivery.ShippingPrice - discount
	} else {
		truckDelivery.DiscountedPrice = truckDelivery.ShippingPrice
	}

	// Obtener la conexión a la base de datos
	db, err := db.GetDBConnection()
	if err != nil {
		http.Error(w, "Error al conectar a la base de datos", http.StatusInternalServerError)
		return
	}

	// Verificar si la tabla "logistics_truck" existe, y si no, crearla
	err = db.AutoMigrate(&models.TruckDelivery{})
	if err != nil {
		http.Error(w, "Error al crear la tabla 'logistics_truck'", http.StatusInternalServerError)
		return
	}

	// Insertar los datos en la tabla "logistics_truck"
	result := db.Create(&truckDelivery)
	if result.Error != nil {
		http.Error(w, fmt.Sprintf("Error al insertar el plan de entrega de logística terrestre: %s", result.Error.Error()), http.StatusInternalServerError)
		return
	}

	fmt.Println("Plan de entrega de logística terrestre creado exitosamente")
}

// UpdateTruckDelivery updates a truck delivery by its ID
// UpdateTruckDelivery updates a truck delivery by its ID
func UpdateTruckDelivery(w http.ResponseWriter, r *http.Request) {
	// Get the truck delivery ID from the request URL
	deliveryID := mux.Vars(r)["id"]
	if deliveryID == "" {
		http.Error(w, "Truck delivery ID not provided", http.StatusBadRequest)
		return
	}

	var updatedDelivery models.TruckDelivery
	err := json.NewDecoder(r.Body).Decode(&updatedDelivery)
	if err != nil {
		http.Error(w, "Error parsing updated truck delivery data", http.StatusBadRequest)
		return
	}

	db, err := db.GetDBConnection()
	if err != nil {
		http.Error(w, "Error connecting to the database", http.StatusInternalServerError)
		return
	}

	// Find the truck delivery by ID
	var delivery models.TruckDelivery
	result := db.First(&delivery, "id = ?", deliveryID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "Truck delivery not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error retrieving truck delivery data", http.StatusInternalServerError)
		}
		return
	}

	// Update the truck delivery with the new data
	result = db.Model(&delivery).Updates(&updatedDelivery)
	if result.Error != nil {
		http.Error(w, "Error updating truck delivery", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(delivery)
}

// DeleteTruckDelivery deletes a truck delivery by its ID
func DeleteTruckDelivery(w http.ResponseWriter, r *http.Request) {
	// Get the truck delivery ID from the request URL
	deliveryID := mux.Vars(r)["id"]
	if deliveryID == "" {
		http.Error(w, "Truck delivery ID not provided", http.StatusBadRequest)
		return
	}

	db, err := db.GetDBConnection()
	if err != nil {
		http.Error(w, "Error connecting to the database", http.StatusInternalServerError)
		return
	}

	// Delete the truck delivery by ID
	result := db.Delete(&models.TruckDelivery{}, deliveryID)
	if result.Error != nil {
		http.Error(w, "Error deleting truck delivery", http.StatusInternalServerError)
		return
	}
	if result.RowsAffected == 0 {
		http.Error(w, "Truck delivery not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
