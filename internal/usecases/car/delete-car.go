package car_use_case

import (
	"context"

	"github.com/cassiusbessa/goCarShop/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c *carUseCase) DeleteCar(ctx context.Context, id string) *utils.CustomError {
	mongoId, _ := primitive.ObjectIDFromHex(id)
	return c.carRepo.DeleteCar(ctx, mongoId)
}
