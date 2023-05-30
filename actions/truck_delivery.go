package actions

import (
	"docucenter-task/db"
	"docucenter-task/internal"
	"docucenter-task/models"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// Handler para crear un plan de entrega de logística marítima
func CreateShipDelivery(w http.ResponseWriter, r *http.Request) {
	// Obtener el token de autenticación de la cabecera de autorización
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Token de autenticación requerido")
		return
	}

	validTk := internal.ValidateToken(tokenString)
	// Validar el token
	if !validTk {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Token de autenticación inválido")
		return
	}

	var shipDelivery models.ShipDelivery
	_ = json.NewDecoder(r.Body).Decode(&shipDelivery)

	// Calcular el descuento
	if shipDelivery.Quantity > 10 {
		discount := shipDelivery.ShippingPrice * 0.03
		shipDelivery.DiscountedPrice = shipDelivery.ShippingPrice - discount
	} else {
		shipDelivery.DiscountedPrice = shipDelivery.ShippingPrice
	}

	validate := validator.New()
	err := validate.Struct(shipDelivery)
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

	// Insertar los datos en la base de datos
	db, err := db.GetDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	// Verificar si la tabla existe y crearla si no
	if !db.Migrator().HasTable(&models.ShipDelivery{}) {
		err = db.AutoMigrate(&models.ShipDelivery{})
		if err != nil {
			log.Fatal(err)
		}
	}

	err = db.Create(&shipDelivery).Error
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Plan de entrega de logística marítima creado exitosamente")
}

// UpdateShipDelivery updates a ship delivery by its ID
func UpdateShipDelivery(w http.ResponseWriter, r *http.Request) {
	// Obtener el token de autenticación de la cabecera de autorización
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Token de autenticación requerido")
		return
	}

	validTk := internal.ValidateToken(tokenString)
	// Validar el token
	if !validTk {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Token de autenticación inválido")
		return
	}

	// Obtener el ID de la entrega de logística marítima desde la URL
	deliveryID := mux.Vars(r)["id"]
	if deliveryID == "" {
		http.Error(w, "ID de entrega de logística marítima no proporcionado", http.StatusBadRequest)
		return
	}

	var updatedDelivery models.ShipDelivery
	err := json.NewDecoder(r.Body).Decode(&updatedDelivery)
	if err != nil {
		http.Error(w, "Error al analizar los datos de la entrega de logística marítima", http.StatusBadRequest)
		return
	}

	// Calcular el descuento
	if updatedDelivery.Quantity > 10 {
		discount := updatedDelivery.ShippingPrice * 0.03
		updatedDelivery.DiscountedPrice = updatedDelivery.ShippingPrice - discount
	} else {
		updatedDelivery.DiscountedPrice = updatedDelivery.ShippingPrice
	}

	validate := validator.New()
	err = validate.Struct(updatedDelivery)
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

	db, err := db.GetDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	// Buscar la entrega de logística marítima por ID
	var delivery models.ShipDelivery
	result := db.First(&delivery, deliveryID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "Entrega de logística marítima no encontrada", http.StatusNotFound)
		} else {
			http.Error(w, "Error al buscar la entrega de logística marítima", http.StatusInternalServerError)
		}
		return
	}

	// Actualizar los datos de la entrega de logística marítima con los nuevos datos
	result = db.Model(&delivery).Updates(updatedDelivery)
	if result.Error != nil {
		http.Error(w, "Error al actualizar la entrega de logística marítima", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Entrega de logística marítima actualizada exitosamente")
}

// DeleteShipDelivery deletes a ship delivery by its ID
func DeleteShipDelivery(w http.ResponseWriter, r *http.Request) {
	// Obtener el token de autenticación de la cabecera de autorización
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Token de autenticación requerido")
		return
	}

	validTk := internal.ValidateToken(tokenString)
	// Validar el token
	if !validTk {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Token de autenticación inválido")
		return
	}

	// Obtener el ID de la entrega de logística marítima desde la URL
	deliveryID := mux.Vars(r)["id"]
	if deliveryID == "" {
		http.Error(w, "ID de entrega de logística marítima no proporcionado", http.StatusBadRequest)
		return
	}

	db, err := db.GetDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	// Buscar la entrega de logística marítima por ID
	var delivery models.ShipDelivery
	result := db.First(&delivery, deliveryID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "Entrega de logística marítima no encontrada", http.StatusNotFound)
		} else {
			http.Error(w, "Error al buscar la entrega de logística marítima", http.StatusInternalServerError)
		}
		return
	}

	// Eliminar la entrega de logística marítima
	result = db.Delete(&delivery)
	if result.Error != nil {
		http.Error(w, "Error al eliminar la entrega de logística marítima", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Entrega de logística marítima eliminada exitosamente")
}
