package routes

import (
	car_controller "github.com/cassiusbessa/goCarShop/internal/controllers/car"
	connection "github.com/cassiusbessa/goCarShop/internal/db"
	"github.com/gorilla/mux"
)

func CarRoutes() *mux.Router {
	carCollection := connection.NewCollection("CarShop", "cars")
	carCtrl := car_controller.New(carCollection)

	router := mux.NewRouter()
	router.HandleFunc("/cars", carCtrl.CreateCar).Methods("POST")
	router.HandleFunc("/cars", carCtrl.GetAllCars).Methods("GET")
	router.HandleFunc("/cars/{id}", carCtrl.GetCar).Methods("GET")
	router.HandleFunc("/cars/{id}", carCtrl.UpdateCar).Methods("PUT")
	router.HandleFunc("/cars/{id}", carCtrl.DeleteCar).Methods("DELETE")

	return router
}
