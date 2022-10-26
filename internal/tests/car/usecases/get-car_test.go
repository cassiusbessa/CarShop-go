package tests

import (
	"context"
	"fmt"
	"testing"

	entity "github.com/cassiusbessa/goCarShop/internal/db/entities"
	"github.com/cassiusbessa/goCarShop/internal/tests/car/mocks"
	car_use_case "github.com/cassiusbessa/goCarShop/internal/usecases/car"
	"github.com/cassiusbessa/goCarShop/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetCarService(t *testing.T) {
	assertions := assert.New(t)
	mockCarUseCase := car_use_case.New(mocks.TestCollection)

	tests := []struct {
		describe string
		arg1     string
		name     string
		want     *entity.ICar
		wantErr  *utils.CustomError
		msg      string
	}{
		{
			describe: "Should get a car",
			arg1:     mocks.Seeds[0].Id.Hex(),
			name:     "GetCar",
			want:     &mocks.Seeds[0],
			wantErr:  nil,
			msg:      "Car got successfully",
		},
		{
			describe: "Should not get a car",
			arg1:     mocks.ErrorCar.Id.Hex(),
			name:     "GetCar",
			want:     nil,
			wantErr:  utils.NewError(400, "Invalid car", "use-case-get-car", nil),
			msg:      "Not got",
		},
	}

	t.Run(tests[0].describe, func(t *testing.T) {
		mocks.Seeder()
		defer mocks.TestCollection.Drop(context.Background())
		got, _ := mockCarUseCase.GetCar(context.Background(), tests[0].arg1)
		fmt.Println(got)
		assertions.Equal(tests[0].want, got, tests[0].msg)
	})

	t.Run(tests[1].describe, func(t *testing.T) {
		_, got := mockCarUseCase.GetCar(context.Background(), tests[1].arg1)
		defer mocks.TestCollection.Drop(context.Background())
		assertions.Error(got, tests[1].msg)
	})
}
