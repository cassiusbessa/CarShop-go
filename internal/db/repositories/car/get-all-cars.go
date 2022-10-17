package car_repo

import (
	"context"
	"net/http"

	entity "github.com/cassiusbessa/goCarShop/internal/db/entities"
	"github.com/cassiusbessa/goCarShop/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c *carRepo) GetAllCars(ctx context.Context) (*[]entity.ICar, *utils.CustomError) {
	var cars []entity.ICar
	cur, err := c.collection.Find(ctx, primitive.M{})
	if err != nil {
		return nil, utils.NewError(http.StatusInternalServerError, "db find error", "repo-get-all-cars", err)
	}
	if err = cur.All(ctx, &cars); err != nil {
		return nil, utils.NewError(http.StatusInternalServerError, "db find error", "repo-get-all-cars", err)
	}
	return &cars, nil
}
