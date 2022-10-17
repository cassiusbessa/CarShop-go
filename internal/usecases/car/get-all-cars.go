package car_use_case

import (
	"context"

	entity "github.com/cassiusbessa/goCarShop/internal/db/entities"
	"github.com/cassiusbessa/goCarShop/internal/utils"
)

func (c *carUseCase) GetAllCars(ctx context.Context) (*[]entity.ICar, *utils.CustomError) {
	return c.carRepo.GetAllCars(ctx)
}
