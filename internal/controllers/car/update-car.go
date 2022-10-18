package car_controller

import (
	"context"
	"encoding/json"
	"net/http"

	entity "github.com/cassiusbessa/goCarShop/internal/db/entities"
	"github.com/cassiusbessa/goCarShop/internal/utils"
	"github.com/gorilla/mux"
)

func (c *carController) UpdateCar(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "application/json")
	param := mux.Vars(req)
	var car *entity.ICar
	err := json.NewDecoder(req.Body).Decode(&car)
	if err != nil {
		err := utils.NewError(http.StatusInternalServerError, "controller-car-update", "Error decoding car", err)
		utils.SendError(err, res)
		return
	}
	car, error := c.carUseCase.UpdateCar(context.Background(), param["id"], *car)
	if err != nil {
		utils.SendError(error, res)
		return
	}
	res.WriteHeader(http.StatusOK)
	defer json.NewEncoder(res).Encode(car)

}
