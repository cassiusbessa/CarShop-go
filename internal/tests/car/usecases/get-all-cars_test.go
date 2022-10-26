package tests

import (
	"context"
	"testing"

	entity "github.com/cassiusbessa/goCarShop/internal/db/entities"
	"github.com/cassiusbessa/goCarShop/internal/tests/car/mocks"
	car_use_case "github.com/cassiusbessa/goCarShop/internal/usecases/car"
	"github.com/cassiusbessa/goCarShop/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetAllCarsService(t *testing.T) {
	assertions := assert.New(t)
	mockCarUseCase := car_use_case.New(mocks.TestCollection)

	tests := []struct {
		describe string
		arg1     *entity.ICar
		name     string
		want     *[]entity.ICar
		wantErr  *utils.CustomError
		msg      string
	}{
		{
			describe: "Should get all cars",
			arg1:     &mocks.Car,
			name:     "GetAllCars",
			want:     &mocks.Seeds,
			wantErr:  nil,
			msg:      "Cars retrieved successfully",
		},
	}
	t.Run(tests[0].describe, func(t *testing.T) {
		mocks.Seeder()
		defer mocks.TestCollection.Drop(context.Background())
		got, _ := mockCarUseCase.GetAllCars(context.Background())
		assertions.Equal(tests[0].want, got, tests[0].msg)
	})
}
