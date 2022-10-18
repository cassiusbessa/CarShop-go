package car_use_case

import (
	"context"

	entity "github.com/cassiusbessa/goCarShop/internal/db/entities"
	"github.com/cassiusbessa/goCarShop/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c *carUseCase) UpdateCar(ctx context.Context, id string, car entity.ICar) (*entity.ICar, *utils.CustomError) {
	mongoId, _ := primitive.ObjectIDFromHex(id)
	return c.carRepo.UpdateCar(ctx, mongoId, car)
}
