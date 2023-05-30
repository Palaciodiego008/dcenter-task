package actions

import (
	"docucenter-task/db"
	"docucenter-task/internal"
	"docucenter-task/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
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
