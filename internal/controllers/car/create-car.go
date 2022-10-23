package car_controller

import (
	"context"
	"encoding/json"
	"net/http"

	entity "github.com/cassiusbessa/goCarShop/internal/db/entities"
	"github.com/cassiusbessa/goCarShop/internal/utils"
)

func (c *carController) CreateCar(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "application/json")
	var car *entity.ICar
	err := json.NewDecoder(req.Body).Decode(&car)
	if err != nil {
		err := utils.NewError(http.StatusInternalServerError, "controller-car-create", "Error decoding car", err)
		utils.SendError(err, res)
		return
	}
	_, error := c.carUseCase.CreateCar(context.Background(), car)
	if err != nil {
		utils.SendError(error, res)
		return
	}
	res.WriteHeader(http.StatusCreated)
	defer json.NewEncoder(res).Encode(car)

}
