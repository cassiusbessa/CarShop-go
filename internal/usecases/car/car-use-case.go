package car_use_case

import (
	"context"

	entity "github.com/cassiusbessa/goCarShop/internal/db/entities"
	car_repo "github.com/cassiusbessa/goCarShop/internal/db/repositories/car"
	"github.com/cassiusbessa/goCarShop/internal/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type ICarUseCase interface {
	CreateCar(ctx context.Context, car entity.ICar) (*entity.ICar, *utils.CustomError)
	GetCar(ctx context.Context, id string) (*entity.ICar, *utils.CustomError)
	GetAllCars(ctx context.Context) (*[]entity.ICar, *utils.CustomError)
	UpdateCar(ctx context.Context, id string, car entity.ICar) (*entity.ICar, *utils.CustomError)
	DeleteCar(ctx context.Context, id string) *utils.CustomError
}

type carUseCase struct {
	carRepo car_repo.ICarRepository
}

func New(collection *mongo.Collection) ICarUseCase {
	return &carUseCase{carRepo: car_repo.New(collection)}
}
