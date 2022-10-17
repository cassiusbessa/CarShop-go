package car_repo

import (
	"context"
	"net/http"

	entity "github.com/cassiusbessa/goCarShop/internal/db/entities"
	"github.com/cassiusbessa/goCarShop/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c *carRepo) GetCar(ctx context.Context, id primitive.ObjectID) (*entity.ICar, *utils.CustomError) {
	var car *entity.ICar
	err := c.collection.FindOne(ctx, primitive.M{"_id": id}).Decode(&car)
	if err != nil {
		return car,
			utils.NewError(http.StatusInternalServerError, "db find error", "repo-get-car", err)
	}
	return car, nil
}
