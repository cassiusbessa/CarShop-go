package car_repo

import (
	"context"

	entity "github.com/cassiusbessa/goCarShop/internal/db/entities"
	"github.com/cassiusbessa/goCarShop/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ICarRepository interface {
	CreateCar(ctx context.Context, car entity.ICar) (*entity.ICar, *utils.CustomError)
	GetCar(ctx context.Context, id primitive.ObjectID) (*entity.ICar, *utils.CustomError)
	GetAllCars(ctx context.Context) (*[]entity.ICar, *utils.CustomError)
}

type carRepo struct {
	collection *mongo.Collection
}

func New(collection *mongo.Collection) ICarRepository {
	return &carRepo{collection: collection}
}
