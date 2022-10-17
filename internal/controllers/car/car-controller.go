package car_controller

import (
	"net/http"

	car_use_case "github.com/cassiusbessa/goCarShop/internal/usecases/car"
	"go.mongodb.org/mongo-driver/mongo"
)

type ICarController interface {
	CreateCar(res http.ResponseWriter, req *http.Request)
	GetCar(res http.ResponseWriter, req *http.Request)
	GetAllCars(res http.ResponseWriter, req *http.Request)
}

type carController struct {
	carUseCase car_use_case.ICarUseCase
}

func New(collection *mongo.Collection) ICarController {
	return &carController{carUseCase: car_use_case.New(collection)}
}
