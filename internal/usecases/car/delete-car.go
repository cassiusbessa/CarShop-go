package car_use_case

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cassiusbessa/goCarShop/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c *carUseCase) DeleteCar(ctx context.Context, id string) *utils.CustomError {
	mongoId, _ := primitive.ObjectIDFromHex(id)
	r, _ := c.carRepo.DeleteCar(ctx, mongoId)
	if r == 0 {
		return utils.NewError(http.StatusNotFound, "car not found", "usecase-delete-car", fmt.Errorf("car not found"))
	}
	return nil
}
