package actions_test

import (
	"docucenter-task/actions"
	"docucenter-task/db"
	"docucenter-task/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// Mock de la interfaz de DB para pruebas
type MockDB struct {
	mock.Mock
}

func (m *MockDB) Model(value interface{}) *gorm.DB {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) Where(query interface{}, args ...interface{}) *gorm.DB {
	argsList := append([]interface{}{query}, args...)
	arg := m.Called(argsList...)
	return arg.Get(0).(*gorm.DB)
}

func (m *MockDB) Find(dest interface{}, conds ...interface{}) *gorm.DB {
	argsList := append([]interface{}{dest}, conds...)
	args := m.Called(argsList...)
	return args.Get(0).(*gorm.DB)
}

func TestGetDeliveries(t *testing.T) {
	// Configurar el mock de la base de datos
	mockDB := new(MockDB)
	db.GetDBConnection()

	// Datos de prueba
	delivery1 := models.TruckDelivery{
		ClientID:         "client1",
		ProductType:      "product1",
		Quantity:         5,
		RegistrationDate: time.Now(),
		DeliveryDate:     time.Now().AddDate(0, 0, 5),
		Warehouse:        "warehouse1",
		ShippingPrice:    100.0,
		DiscountedPrice:  95.0,
		VehiclePlate:     "ABC123",
		GuideNumber:      "1234567890",
	}
	delivery2 := models.TruckDelivery{
		ClientID:         "client2",
		ProductType:      "product2",
		Quantity:         3,
		RegistrationDate: time.Now(),
		DeliveryDate:     time.Now().AddDate(0, 0, 3),
		Warehouse:        "warehouse2",
		ShippingPrice:    200.0,
		DiscountedPrice:  200.0,
		VehiclePlate:     "XYZ789",
		GuideNumber:      "0987654321",
	}
	deliveries := []models.TruckDelivery{delivery1, delivery2}

	// Mock de la consulta de GORM
	mockQuery := new(MockDB)
	mockQuery.On("Find", mock.Anything, mock.Anything).Return(mockQuery)
	mockQuery.On("Error").Return(nil)
	mockQuery.On("Scan", mock.Anything).Run(func(args mock.Arguments) {
		dest := args.Get(0).(*[]models.TruckDelivery)
		*dest = deliveries
	}).Return(mockQuery)

	// Configurar el mock de la base de datos
	mockDB.On("Model", mock.Anything).Return(mockQuery)

	// Crear una solicitud de prueba
	reqBody := ""
	req, err := http.NewRequest("GET", "/deliveries?client_id=client1&product_type=product1", strings.NewReader(reqBody))
	assert.NoError(t, err)

	// Crear un ResponseRecorder (implementa http.ResponseWriter) para obtener la respuesta
	rr := httptest.NewRecorder()

	// Llamar al controlador
	handler := http.HandlerFunc(actions.GetDeliveries)
	handler.ServeHTTP(rr, req)

	// Verificar el código de estado de la respuesta
	assert.Equal(t, http.StatusOK, rr.Code)

	// Verificar el tipo de contenido de la respuesta
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	// Verificar los resultados de la respuesta
	var responseDeliveries []models.TruckDelivery
	err = json.Unmarshal(rr.Body.Bytes(), &responseDeliveries)
	assert.NoError(t, err)
	assert.Equal(t, deliveries, responseDeliveries)

	// Verificar las llamadas al mock de la base de datos
	mockDB.AssertExpectations(t)
	mockQuery.AssertExpectations(t)
}
