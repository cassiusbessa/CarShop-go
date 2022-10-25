package car_use_case

import (
	"context"
	"net/http"
	"reflect"

	entity "github.com/cassiusbessa/goCarShop/internal/db/entities"
	"github.com/cassiusbessa/goCarShop/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c *carUseCase) UpdateCar(ctx context.Context, id string, car entity.ICar) (*entity.ICar, *utils.CustomError) {
	mongoId, _ := primitive.ObjectIDFromHex(id)
	err := validateUpdate(car)
	if err != nil {
		return &entity.ICar{}, utils.NewError(http.StatusBadRequest, err.Error(), "use-case-update-car", err)
	}
	return c.carRepo.UpdateCar(ctx, mongoId, car)
}

func validateUpdate(car entity.ICar) *utils.CustomError {
	v := reflect.ValueOf(car)
	keys := reflect.VisibleFields(v.Type())
	var validateKeys []string
	for _, key := range keys {
		if !(v.FieldByName(key.Name).IsZero()) {
			validateKeys = append(validateKeys, key.Name)
		}
	}
	err := car.UpdateValidate(validateKeys...)
	if err != nil {
		return utils.NewError(http.StatusBadRequest, err.Error(), "use-case-update-car", err)
	}
	return nil
}
