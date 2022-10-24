package car_repo

import (
	"context"
	"net/http"

	"github.com/cassiusbessa/goCarShop/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c *carRepo) DeleteCar(ctx context.Context, id primitive.ObjectID) (int64, *utils.CustomError) {
	r, err := c.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return 0, utils.NewError(http.StatusInternalServerError, "db delete error", "repo-delete-car", err)
	}
	return r.DeletedCount, nil
}
