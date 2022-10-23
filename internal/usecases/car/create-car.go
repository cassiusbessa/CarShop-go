package car_use_case

import (
	"context"
	"net/http"

	entity "github.com/cassiusbessa/goCarShop/internal/db/entities"
	"github.com/cassiusbessa/goCarShop/internal/utils"
)

func (c *carUseCase) CreateCar(ctx context.Context, car *entity.ICar) (*entity.ICar, *utils.CustomError) {
	err := car.Validate()
	if err != nil {
		return &entity.ICar{}, utils.NewError(http.StatusBadRequest, err.Error(), "use-case-create-car", err)
	}
	return c.carRepo.CreateCar(ctx, car)
}
