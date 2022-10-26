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

func TestCreateCarService(t *testing.T) {
	assertions := assert.New(t)
	mockCarUseCase := car_use_case.New(mocks.TestCollection)

	tests := []struct {
		describe string
		arg1     *entity.ICar
		name     string
		want     *entity.ICar
		wantErr  *utils.CustomError
		msg      string
	}{
		{
			describe: "Should create a car",
			arg1:     &mocks.Car,
			name:     "CreateCar",
			want:     &mocks.Car,
			wantErr:  nil,
			msg:      "Car created successfully",
		},

		{
			describe: "Should not create a car",
			arg1:     &mocks.ErrorCar,
			name:     "CreateCar",
			want:     nil,
			wantErr:  utils.NewError(400, "Invalid car", "use-case-create-car", nil),
			msg:      "Not created",
		},
	}
	t.Run(tests[0].describe, func(t *testing.T) {
		got, _ := mockCarUseCase.CreateCar(context.Background(), tests[0].arg1)
		defer mocks.TestCollection.Drop(context.Background())
		assertions.Equal(tests[0].want, got, tests[0].msg)
	})

	t.Run(tests[1].describe, func(t *testing.T) {
		_, got := mockCarUseCase.CreateCar(context.Background(), tests[1].arg1)
		defer mocks.TestCollection.Drop(context.Background())
		assertions.Error(got, tests[1].msg)
	})
}
