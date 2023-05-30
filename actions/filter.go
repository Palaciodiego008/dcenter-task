package actions

import (
	"net/http"
)

// Handler para consultar los planes de entrega con filtros
func GetDeliveries(w http.ResponseWriter, r *http.Request) {
	// // Obtener los parámetros de consulta
	// queryParams := r.URL.Query()

	// // Construir la consulta SQL base
	// sqlQuery := "SELECT * FROM logistics_truck WHERE 1 = 1"

	// // Agregar filtros según los parámetros de consulta
	// if clientID := queryParams.Get("client_id"); clientID != "" {
	// 	sqlQuery += fmt.Sprintf(" AND client_id = %s", clientID)
	// }
	// if productType := queryParams.Get("product_type"); productType != "" {
	// 	sqlQuery += fmt.Sprintf(" AND product_type = '%s'", productType)
	// }
	// // Agrega más filtros según sea necesario

	// // Ejecutar la consulta en la base de datos
	// db := db.GetDBConnection()
	// defer db.Close()

	// rows, err := db.Query(sqlQuery)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()

	// // Iterar sobre los resultados y enviar la respuesta
	// deliveries := []models.TruckDelivery{}
	// for rows.Next() {
	// 	var delivery models.TruckDelivery
	// 	err := rows.Scan(
	// 		&delivery.ClientID, &delivery.ProductType, &delivery.Quantity, &delivery.RegistrationDate,
	// 		&delivery.DeliveryDate, &delivery.Warehouse, &delivery.ShippingPrice, &delivery.DiscountedPrice,
	// 		&delivery.VehiclePlate, &delivery.GuideNumber)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	deliveries = append(deliveries, delivery)
	// }

	// // Convertir los resultados a JSON y enviar la respuesta
	// response, err := json.Marshal(deliveries)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write(response)
}
