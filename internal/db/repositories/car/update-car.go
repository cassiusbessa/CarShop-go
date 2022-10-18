package car_repo

import (
	"context"
	"net/http"

	entity "github.com/cassiusbessa/goCarShop/internal/db/entities"
	"github.com/cassiusbessa/goCarShop/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c *carRepo) UpdateCar(ctx context.Context, id primitive.ObjectID, car entity.ICar) (*entity.ICar, *utils.CustomError) {
	r := c.collection.FindOneAndUpdate(context.Background(), bson.M{"_id": id}, bson.M{"$set": car})
	if r.Err() != nil {
		return &entity.ICar{},
			utils.NewError(http.StatusInternalServerError, "db update error", "repo-update-car", r.Err())
	}
	r.Decode(&car)
	return &car, nil
}
