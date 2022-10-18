package car_controller

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/cassiusbessa/goCarShop/internal/utils"
	"github.com/gorilla/mux"
)

func (c *carController) DeleteCar(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "application/json")
	param := mux.Vars(req)
	err := c.carUseCase.DeleteCar(context.Background(), param["id"])
	if err != nil {
		utils.SendError(err, res)
		return
	}
	res.WriteHeader(http.StatusOK)
	defer json.NewEncoder(res).Encode("Car deleted")
}
