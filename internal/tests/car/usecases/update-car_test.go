package tests

import (
	"context"
	"testing"

	entity "github.com/cassiusbessa/goCarShop/internal/db/entities"
	"github.com/cassiusbessa/goCarShop/internal/tests/car/mocks"
	car_use_case "github.com/cassiusbessa/goCarShop/internal/usecases/car"
	"github.com/cassiusbessa/goCarShop/internal/utils"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUpdateCarService(t *testing.T) {
	assertions := assert.New(t)
	mockCarUseCase := car_use_case.New(mocks.TestCollection)
	tests := []struct {
		describe string
		arg1     string
		arg2     entity.ICar
		name     string
		want     *entity.ICar
		wantErr  *utils.CustomError
		msg      string
	}{
		{
			describe: "Should update a car",
			arg1:     mocks.Seeds[0].Id.Hex(),
			arg2:     mocks.UpdateCar,
			name:     "UpdateCar",
			want:     &mocks.Seeds[0],
			wantErr:  nil,
			msg:      "Car updated successfully",
		},

		{
			describe: "Should not update a car",
			arg1:     primitive.NewObjectID().Hex(),
			arg2:     mocks.UpdateCar,
			name:     "UpdateCar",
			want:     nil,
			wantErr:  utils.NewError(400, "Invalid car", "use-case-update-car", nil),
			msg:      "Not updated",
		},
	}
	t.Run(tests[0].describe, func(t *testing.T) {
		mocks.Seeder()
		defer mocks.TestCollection.Drop(context.Background())
		mockCarUseCase.UpdateCar(context.Background(), tests[0].arg1, tests[0].arg2)
		got, _ := mockCarUseCase.GetCar(context.Background(), tests[0].arg1)
		assertions.NotEqual(tests[0].want, got, tests[0].msg)
	})

	t.Run(tests[1].describe, func(t *testing.T) {
		_, got := mockCarUseCase.UpdateCar(context.Background(), tests[1].arg1, tests[1].arg2)
		defer mocks.TestCollection.Drop(context.Background())
		assertions.Error(got, tests[1].msg)
	})
}
