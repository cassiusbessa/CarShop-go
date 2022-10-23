package car_repo

import (
	"context"
	"net/http"

	entity "github.com/cassiusbessa/goCarShop/internal/db/entities"
	"github.com/cassiusbessa/goCarShop/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c *carRepo) CreateCar(ctx context.Context, car *entity.ICar) (*entity.ICar, *utils.CustomError) {
	car.Id = primitive.NewObjectID()
	_, err := c.collection.InsertOne(ctx, car)
	if err != nil {
		return &entity.ICar{},
			utils.NewError(http.StatusInternalServerError, "db insert error", "repo-create-car", err)
	}
	return car, nil
}
