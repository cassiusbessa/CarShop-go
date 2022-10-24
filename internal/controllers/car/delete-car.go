package car_controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cassiusbessa/goCarShop/internal/utils"
	"github.com/gorilla/mux"
)

func (c *carController) DeleteCar(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "application/json")
	param := mux.Vars(req)
	err := c.carUseCase.DeleteCar(context.Background(), param["id"])
	fmt.Println("entrei aqui >>>>>>>>", err)
	if err != nil {
		fmt.Println("entrei aqui controler", err)
		utils.SendError(err, res)
		return
	}
	res.WriteHeader(http.StatusOK)
	defer json.NewEncoder(res).Encode("successfully deleted car")
}
