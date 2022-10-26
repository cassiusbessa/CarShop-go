package car_use_case

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"

	entity "github.com/cassiusbessa/goCarShop/internal/db/entities"
	"github.com/cassiusbessa/goCarShop/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c *carUseCase) UpdateCar(ctx context.Context, id string, car entity.ICar) (*entity.ICar, *utils.CustomError) {
	mongoId, _ := primitive.ObjectIDFromHex(id)
	err := validateUpdateCar(car)
	if err != nil {
		return &entity.ICar{}, utils.NewError(http.StatusBadRequest, err.Error(), "use-case-update-car", err)
	}
	return c.carRepo.UpdateCar(ctx, mongoId, car)
}

func validateUpdateCar(car entity.ICar) error {
	var inInterface map[string]interface{}
	carrec, _ := json.Marshal(car)
	json.Unmarshal(carrec, &inInterface)
	k := reflect.ValueOf(inInterface).MapKeys()
	keys := make([]string, len(k))
	for i, a := range k {
		keys[i] = a.String()
	}
	return car.UpdateValidate(keys...)
}
