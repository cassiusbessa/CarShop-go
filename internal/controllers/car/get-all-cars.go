package car_controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/cassiusbessa/goCarShop/internal/utils"
)

func (c *carController) GetAllCars(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "application/json")
	cars, err := c.carUseCase.GetAllCars(context.Background())
	if err != nil {
		utils.SendError(err, res)
		return
	}
	res.WriteHeader(http.StatusOK)
	defer json.NewEncoder(res).Encode(cars)
}
