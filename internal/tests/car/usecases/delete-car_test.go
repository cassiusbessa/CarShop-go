package tests

import (
	"context"
	"net/http"
	"testing"

	"github.com/cassiusbessa/goCarShop/internal/tests/car/mocks"
	car_use_case "github.com/cassiusbessa/goCarShop/internal/usecases/car"
	"github.com/cassiusbessa/goCarShop/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestDeleteCarService(t *testing.T) {
	assertions := assert.New(t)
	mockCarUseCase := car_use_case.New(mocks.TestCollection)

	tests := []struct {
		describe string
		arg1     string
		name     string
		want     *utils.CustomError
		wantErr  *utils.CustomError
		msg      string
	}{
		{
			describe: "Should delete a car",
			arg1:     mocks.Seeds[0].Id.Hex(),
			name:     "DeleteCar",
			want:     nil,
			wantErr:  nil,
			msg:      "Car deleted successfully",
		},

		{
			describe: "Should not delete a car",
			arg1:     mocks.ErrorCar.Id.Hex(),
			name:     "DeleteCar",
			want:     nil,
			wantErr:  utils.NewError(http.StatusBadRequest, "Invalid car", "use-case-delete-car", nil),
			msg:      "Not deleted",
		},
	}
	t.Run(tests[0].describe, func(t *testing.T) {
		mocks.Seeder()
		got := mockCarUseCase.DeleteCar(context.Background(), tests[0].arg1)
		defer mocks.TestCollection.Drop(context.Background())
		assertions.Equal(tests[0].want, got, tests[0].msg)
	})

	t.Run(tests[1].describe, func(t *testing.T) {
		mocks.Seeder()
		got := mockCarUseCase.DeleteCar(context.Background(), tests[1].arg1)
		defer mocks.TestCollection.Drop(context.Background())
		assertions.Error(got, tests[1].msg)
	})
}
