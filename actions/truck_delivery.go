package actions

import (
	"docucenter-task/db"
	"docucenter-task/internal"
	"docucenter-task/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator"
)

// Handler para crear un plan de entrega de logística marítima

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
