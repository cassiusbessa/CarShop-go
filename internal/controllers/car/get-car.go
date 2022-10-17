package car_controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/cassiusbessa/goCarShop/internal/utils"
	"github.com/gorilla/mux"
)

func (c *carController) GetCar(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "application/json")
	param := mux.Vars(req)
	car, err := c.carUseCase.GetCar(context.Background(), param["id"])
	if err != nil {
		utils.SendError(err, res)
		return
	}
	res.WriteHeader(http.StatusOK)
	defer json.NewEncoder(res).Encode(car)
}
