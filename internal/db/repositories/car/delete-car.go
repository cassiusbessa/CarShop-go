package car_repo

import (
	"context"
	"net/http"

	"github.com/cassiusbessa/goCarShop/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (c *carRepo) DeleteCar(ctx context.Context, id primitive.ObjectID) *utils.CustomError {
	_, err := c.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return utils.NewError(http.StatusInternalServerError, "db delete error", "repo-delete-car", err)
	}
	return nil
}
